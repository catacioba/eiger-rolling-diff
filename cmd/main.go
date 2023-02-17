package main

import (
	"flag"
	"fmt"
	"github.com/catacioba/rdiff"
	"os"
)

func main() {

	chunkSize := flag.Int("chunkSize", 1024, "block chunk size")

	flag.Parse()

	oldFilename := flag.Args()[0]
	newFilename := flag.Args()[1]

	oldFile, err := os.Open(oldFilename)
	if err != nil {
		panic(err)
	}

	newFile, err := os.Open(newFilename)
	if err != nil {
		panic(err)
	}

	chunks, err := rdiff.Signature(oldFile, *chunkSize)
	if err != nil {
		panic(err)
	}

	fmt.Println("blocks:")
	for _, chunk := range chunks {
		fmt.Printf("%v\n", chunk)
	}

	diff, err := rdiff.DiffFiles(newFile, chunks)
	if err != nil {
		panic(err)
	}

	fmt.Println("matches:")
	for idx := range diff.Matches {
		fmt.Printf("-> %v\n", diff.Matches[idx])
	}

	fmt.Println("additions:")
	for idx := range diff.Additions {
		fmt.Printf("-> %v %c\n", diff.Additions[idx], diff.Additions[idx].Data)
	}
}
