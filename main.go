package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("text.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	blocks := MarkdownToBlocks(string(data))

	for _, b := range blocks {
		block := BlockParser(b)
		fmt.Println(block)
		fmt.Println("---")
	}
}
