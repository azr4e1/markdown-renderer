package main

import (
	"fmt"
	"strings"
)

const (
	BLOCKDELIMITER   = "\n\n"
	HEADERPREFIX     = "#"
	BREAKDELIMITER   = "---"
	CODEDELIMITER    = "```"
	QUOTEPREFIX1     = "> "
	QUOTEPREFIX2     = "  "
	QUOTEPREFIX3     = "\t"
	UNORDEREDPREFIX1 = "* "
	UNORDEREDPREFIX2 = "- "
	ORDEREDPREFIX    = "1. "
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
	if level, isH := isHeader(block); isH {
		return headerify(block, level)
	} else if isBreak(block) {
		return Break(true)
	} else if isCode(block) {
		return codeify(block)
	} else if del, isQ := isQuote(block); isQ {
		return quoteify(block, del)
	} else if isOrderedList(block) {
		return olistify(block)
	} else if isUnorderedList(block) {
		return ulistify(block)
	} else if isTable(block) {
		return tableify(block)
	} else {
		return Paragraph(LineParser(block))
	}
}

func isHeader(block string) (int, bool) {
	if len(strings.Split(block, "\n")) != 1 {
		return 0, false
	}
	var level int
	isH := false
	for ; strings.HasPrefix(block, HEADERPREFIX); level++ {
		isH = true
		block = strings.TrimPrefix(block, HEADERPREFIX)
	}
	// test if there is at least a single space
	isH = strings.HasPrefix(block, " ")

	return level, isH
}

func isBreak(block string) bool {
	return block == BREAKDELIMITER
}

func isCode(block string) bool {
	return strings.HasPrefix(block, CODEDELIMITER) && strings.HasSuffix(block, CODEDELIMITER)
}

func isQuote(block string) (string, bool) {
	// it either starts with >, a tab or a similar amount of whitespce
	lines := strings.Split(block, "\n")
	if len(lines) == 0 {
		return "", false
	}
	firstLine := lines[0]
	if len(firstLine) < 2 {
		return "", false
	}
	var delimiter string
	if strings.HasPrefix(firstLine, QUOTEPREFIX1) {
		delimiter = QUOTEPREFIX1
	} else if strings.HasPrefix(firstLine, QUOTEPREFIX2) {
		delimiter = QUOTEPREFIX2
	} else if strings.HasPrefix(firstLine, QUOTEPREFIX3) {
		delimiter = QUOTEPREFIX3
	} else {
		return "", false
	}
	for _, l := range lines {
		if !strings.HasPrefix(l, delimiter) {
			return "", false
		}
	}
	return delimiter, true
}

func isUnorderedList(block string) bool {
	lines := strings.Split(block, "\n")
	if len(lines) == 0 {
		return false
	}
	for _, l := range lines {
		if !strings.HasPrefix(l, UNORDEREDPREFIX1) && !strings.HasPrefix(l, UNORDEREDPREFIX2) {
			return false
		}
	}

	return true
}

func isOrderedList(block string) bool {
	lines := strings.Split(block, "\n")
	if len(lines) == 0 {
		return false
	}
	counter := 1
	for _, l := range lines {
		countedPrefix := fmt.Sprintf("%d. ", counter)
		if !strings.HasPrefix(l, ORDEREDPREFIX) && !strings.HasPrefix(l, countedPrefix) {
			return false
		}
		counter++
	}
	return true
}

func isTable(block string) bool {
	return false
}

func headerify(block string, level int) Header {
	prefix := ""
	for i := 0; i < level; i++ {
		prefix += HEADERPREFIX
	}
	prefix += " "
	content := strings.TrimPrefix(block, prefix)

	return Header{LineParser(content), level}
}

func codeify(block string) Code {
	return Code(strings.Trim(block, CODEDELIMITER))
}

func quoteify(block, delimiter string) Quote {
	lines := strings.Split(block, "\n")
	if len(lines) == 0 {
		return nil
	}
	newLines := []string{}

	for _, l := range lines {
		newLine := strings.TrimPrefix(l, delimiter)
		newLines = append(newLines, newLine)
	}
	content := strings.Join(newLines, "\n")

	return Quote(LineParser(content))
}

// subtasks to implement
func ulistify(block string) UnorderedList {
	lines := strings.Split(block, "\n")
	if len(lines) == 0 {
		return nil
	}
	newLines := []UnorderedItem{}
	for _, l := range lines {
		newLine := strings.TrimPrefix(l, UNORDEREDPREFIX1)
		if newLine == l {
			newLine = strings.TrimPrefix(l, UNORDEREDPREFIX2)
		}
		newItem := UnorderedItem(LineParser(newLine))
		newLines = append(newLines, newItem)
	}

	return newLines
}

func olistify(block string) OrderedList {
	lines := strings.Split(block, "\n")
	if len(lines) == 0 {
		return nil
	}
	newLines := []OrderedItem{}
	counter := 1
	for _, l := range lines {
		countedPrefix := fmt.Sprintf("%d. ", counter)
		newLine := strings.TrimPrefix(l, ORDEREDPREFIX)
		if newLine == l {
			newLine = strings.TrimPrefix(l, countedPrefix)
		}
		newItem := OrderedItem(LineParser(newLine))
		newLines = append(newLines, newItem)
		counter++
	}

	return newLines
}

func tableify(block string) Table {
	return Table{}
}
