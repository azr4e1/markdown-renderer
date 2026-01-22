package main

import (
	"strings"
)

const (
	BLOCKDELIMITER = "\n\n"
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
