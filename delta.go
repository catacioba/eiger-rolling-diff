package rdiff

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type Delta struct {
	Matches   []Match
	Additions []Addition
}

type Match struct {
	newIdx   int
	oldIdx   int
	blockLen int
}

type Addition struct {
	NewIdx int
	Data   byte
}

func DiffFiles(file io.Reader, chunks []ChunkMetadata) (Delta, error) {
	reader := bufio.NewReader(file)

	chunkMap := make(map[uint16]*ChunkMetadata)
	for idx := range chunks {
		chunk := chunks[idx]
		chunkMap[chunk.weakChecksum] = &chunk
	}

	if len(chunks) == 0 {
		return Delta{}, fmt.Errorf("empty signature data")
	}

	chunkSize := chunks[0].blockLen

	rollSum := NewPolynomialRollingHash()
	ringBuf := NewArrayRingBuffer(chunkSize)

	matches := make([]Match, 0)
	additions := make([]Addition, 0)

	position := 0
	for {
		b, err := reader.ReadByte()
		if err != io.EOF && err != nil {
			return Delta{}, err
		}

		if ringBuf.Len() == chunkSize {
			oldest := ringBuf.Pop()

			addition := Addition{
				NewIdx: position - chunkSize,
				Data:   oldest,
			}
			additions = append(additions, addition)

			rollSum.RotatePush(b, oldest)
		} else {
			rollSum.Push(b)
		}

		ringBuf.Push(b)

		if ringBuf.Len() < chunkSize {
			position += 1
			continue
		}

		weakSum := rollSum.CheckSum()

		if chunk, found := chunkMap[weakSum]; found {
			strong := StrongCheckSum(ringBuf.Data())

			if bytes.Equal(strong[:], chunk.strongChecksum[:]) {
				match := Match{
					newIdx:   position - chunkSize + 1,
					oldIdx:   chunk.start,
					blockLen: chunk.blockLen,
				}
				matches = append(matches, match)
				ringBuf.Clear()
				rollSum.Reset()
			}
		}

		position += 1

		if err == io.EOF {
			break
		}
	}

	for ringBuf.Len() > 0 {
		addition := Addition{
			NewIdx: position - chunkSize,
			Data:   ringBuf.Pop(),
		}
		additions = append(additions, addition)
		position++
	}

	return Delta{
		Matches:   matches,
		Additions: additions,
	}, nil
}
