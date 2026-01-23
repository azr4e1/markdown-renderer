package main

import (
	"strings"
)

const (
	BLOCKDELIMITER = "\n\n"
	BREAKDELIMITER = "---"
	CODEDELIMITER  = "```"
	QUOTEDELIMITER = ">"
)

func MarkdownToBlocks(markdown string) []string {
	blocks := strings.Split(markdown, BLOCKDELIMITER)
	cleanBlocks := make([]string, 0, len(blocks))
	for _, b := range blocks {
		cleanBlock := strings.TrimSpace(b)
		if len(cleanBlock) == 0 {
			continue
		}
		cleanBlocks = append(cleanBlocks, cleanBlock)
	}

	return cleanBlocks
}

func BlockParser(block string) Block {
	switch {
	case isBreak(block):
		return Break(true)
	case isCode(block):
		return codeify(block)
	case isQuote(block):
		return quoteify(block)
	case isOrderedList(block):
		return listify(block, 0)
	case isUnorderedList(block):
		return listify(block, 1)
	case isTable(block):
		return tableify(block)
	default:
		return Paragraph{Text(block)}
	}
}

func isBreak(block string) bool {
	return block == BREAKDELIMITER
}

func isCode(block string) bool {
	return strings.HasPrefix(block, CODEDELIMITER) && strings.HasSuffix(block, CODEDELIMITER)
}

func codeify(block string) Code {
	return Code(strings.Trim(block, CODEDELIMITER))
}

func isQuote
