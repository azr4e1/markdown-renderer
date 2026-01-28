package main

import (
	"fmt"
	"io"
	"os"

	mdr "github.com/azr4e1/markdown-renderer"
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

	html := mdr.MarkdownToHTML(string(data))

	fmt.Println(html.HTMLRender())

}
