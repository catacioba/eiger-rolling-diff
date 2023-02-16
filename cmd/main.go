package main

import (
	"fmt"
	rdiff "github.com/catacioba/eiger-rolling-diff"
	"os"
)

func main() {
	oldFilename := os.Args[1]
	//newFilename := os.Args[2]

	chunks, err := rdiff.Signature(oldFilename, 1024)
	if err != nil {
		panic(err)
	}
	for _, chunk := range chunks {
		fmt.Printf("%v\n", chunk)
	}
}
