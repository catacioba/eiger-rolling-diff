package rdiff

import (
	"bufio"
	"io"
)

type ChunkMetadata struct {
	start          int
	blockLen       int
	weakChecksum   uint16
	strongChecksum [16]byte
}

func Signature(file io.Reader, chunkSize int) ([]ChunkMetadata, error) {
	reader := bufio.NewReader(file)
	buf := make([]byte, chunkSize)

	chunks := make([]ChunkMetadata, 0)

	position := 0
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		rollSum := NewPolynomialRollingHash()

		for _, b := range buf[0:n] {
			rollSum.Push(b)
		}

		strong := StrongCheckSum(buf[0:n])

		chunk := ChunkMetadata{
			start:          position,
			blockLen:       n,
			weakChecksum:   rollSum.CheckSum(),
			strongChecksum: strong,
		}
		chunks = append(chunks, chunk)

		position += n
	}

	return chunks, nil
}
