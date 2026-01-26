package main

import (
	"fmt"
	"strings"
)

const (
	BOLDDELIMITER1      = "**"
	BOLDDELIMITER2      = "__"
	ITALICDELIMITER1    = "*"
	ITALICDELIMITER2    = "_"
	UNDERLINEDELIMITER  = "~"
	INLINECODEDELIMITER = "`"
	CROSSEDDELIMITER    = "-"
)

const (
	PLAIN = iota
	BOLD
	ITALIC
	UNDERLINE
	INLINECODE
	CROSSED
)

var DELIMITERS map[int][]string = map[int][]string{
	BOLD:       []string{BOLDDELIMITER1, BOLDDELIMITER2},
	ITALIC:     []string{ITALICDELIMITER1, ITALICDELIMITER2},
	UNDERLINE:  []string{UNDERLINEDELIMITER},
	INLINECODE: []string{INLINECODEDELIMITER},
	CROSSED:    []string{CROSSEDDELIMITER},
}

// func LineParser(nodes []Text) []Text {
// 	newNodes := []Text{}
// 	for _, n := range nodes {
// 		switch node := n.(type) {
// 		case Plain:
// 			fmt.Println(node)
// 			return nil
// 		default:
// 			newNodes = append(newNodes, n)
// 			continue
// 		}
// 	}
// 	return nil
// }

func LineParser(line string) []Text {
	if len(line) < 2 {
		return []Text{Plain(line)}
	}

	window := line[:1]
	inside := false
	delimiter := ""
	inlineType := PLAIN
	text := ""
	nodes := []Text{}

	for i := 1; i < len(line); i++ {
		current := string(line[i])
		fmt.Println(current)
		window = window[len(window)-1:] + current
		if !inside {
			switch {
			case strings.HasPrefix(window, BOLDDELIMITER1):
				delimiter = BOLDDELIMITER1
				inlineType = BOLD
				inside = true
			case strings.HasPrefix(window, BOLDDELIMITER2):
				delimiter = BOLDDELIMITER2
				inlineType = BOLD
				inside = true
			case strings.HasPrefix(window, ITALICDELIMITER1):
				delimiter = ITALICDELIMITER1
				inlineType = ITALIC
				inside = true
			case strings.HasPrefix(window, ITALICDELIMITER2):
				delimiter = ITALICDELIMITER2
				inlineType = ITALIC
				inside = true
			case strings.HasPrefix(window, UNDERLINEDELIMITER):
				delimiter = UNDERLINEDELIMITER
				inlineType = UNDERLINE
				inside = true
			case strings.HasPrefix(window, INLINECODEDELIMITER):
				delimiter = INLINECODEDELIMITER
				inlineType = INLINECODE
				inside = true
			case strings.HasPrefix(window, CROSSEDDELIMITER):
				delimiter = CROSSEDDELIMITER
				inlineType = CROSSED
				inside = true
			default:
				delimiter = ""
				inside = false
				if i == 1 {
					text += window[:1]
				}
			}
			if inside {
				text = text[:len(text)-1]
				nodes = append(nodes, Plain(text))
				text = strings.TrimPrefix(window, delimiter)
				i++
				if i < len(line) {
					window = string(line[i])
				}
				continue
			}
		} else {
			if strings.HasSuffix(window, delimiter) {
				text += strings.TrimSuffix(window, delimiter)
				nodes = append(nodes, SetType(text, inlineType))
				text = ""
				inlineType = PLAIN
				delimiter = ""
				inside = false
			}
		}
		text += current
	}
	if len(text) != 0 {
		nodes = append(nodes, SetType(text, inlineType))
	}
	return nodes
}

func SetType(text string, TYPE int) Text {
	var typedText Text
	switch TYPE {
	case BOLD:
		typedText = Bold(text)
	case ITALIC:
		typedText = Italic(text)
	case UNDERLINE:
		typedText = Underline(text)
	case INLINECODE:
		typedText = InlineCode(text)
	case CROSSED:
		typedText = Crossed(text)
	default:
		typedText = Plain(text)
	}

	return typedText
}
