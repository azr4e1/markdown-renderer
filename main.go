package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("index.md")
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

// func main() {
// 	text := "![This is an image](hello0.png) *Ci ao c om*-ok -![This is an image](hello1.png) ![This is an image](hello2.png) **e va**"
// 	nodes := NodeParser([]Text{Plain(text)})

// 	fmt.Println(text)
// 	fmt.Println(nodes)
// }
