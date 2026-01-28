package markdownrenderer

func MarkdownToHTML(content string) HTMLNode {
	blocks := MarkdownToBlocks(content)
	htmlNodes := []HTMLNode{}
	for _, b := range blocks {
		blockNode := BlockParser(b)
		htmlNodes = append(htmlNodes, blockNode.ToHTML())
	}

	return HTMLDiv(htmlNodes)
}
