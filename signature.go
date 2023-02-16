package rdiff

import (
	"bufio"
	"io"
	"os"
)

type ChunkMetadata struct {
	start          int
	size           int
	weakChecksum   uint16
	strongChecksum [16]byte
}

func Signature(filename string, chunkSize int) ([]ChunkMetadata, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return SignatureReader(file, chunkSize)
}

func SignatureReader(file io.Reader, chunkSize int) ([]ChunkMetadata, error) {
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
			size:           n,
			weakChecksum:   rollSum.ChuckSum(),
			strongChecksum: strong,
		}
		chunks = append(chunks, chunk)

		position += n
	}

	return chunks, nil
}
