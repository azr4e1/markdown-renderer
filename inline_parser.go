package main

import (
	"regexp"
	"strings"
)

const (
	BOLDDELIMITER1      = "**"
	BOLDDELIMITER2      = "__"
	ITALICDELIMITER1    = "*"
	ITALICDELIMITER2    = "_"
	UNDERLINEDELIMITER  = "-"
	INLINECODEDELIMITER = "`"
	CROSSEDDELIMITER    = "~"
	IMAGEREGEX          = "!\\[(.*?)\\]\\((.*?)\\)"
	LINKREGEX           = "\\[(.*?)\\]\\((.*?)\\)"
)

const (
	PLAIN = iota
	BOLD
	ITALIC
	UNDERLINE
	INLINECODE
	CROSSED
)

var ImagePattern = regexp.MustCompile(IMAGEREGEX)
var HyperlinkPattern = regexp.MustCompile(LINKREGEX)

func SimpleParser(line string) []Node {
	if len(line) < 2 {
		return []Node{Plain(line)}
	}

	window := line[:1]
	inside := false
	delimiter := ""
	inlineType := PLAIN
	text := window
	nodes := []Node{}

	for i := 1; i < len(line); i++ {
		current := string(line[i])
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
			}
			if inside {
				text = strings.TrimSuffix(text+current, window)
				if len(text) > 0 {
					nodes = append(nodes, Plain(text))
				}
				text = strings.TrimPrefix(window, delimiter)
				if i < len(line) {
					window = string(line[i])
				}
				continue
			}
		} else {
			if strings.HasSuffix(window, delimiter) {
				text = strings.TrimSuffix(text+current, delimiter)
				nodes = append(nodes, SetType(text, inlineType))
				text = ""
				inlineType = PLAIN
				delimiter = ""
				inside = false
				i++
				if i < len(line) {
					window = string(line[i])
					text = window
				}
				continue
			}
		}
		text += current
	}
	if len(text) != 0 {
		nodes = append(nodes, SetType(text, inlineType))
	}
	return nodes
}

func ImageParser(line string) []Node {
	// found := ImagePattern.FindAllString(line, -1)
	found := ImagePattern.FindAllStringSubmatch(line, -1)
	textNodes := ImagePattern.Split(line, -1)
	nodes := []Node{}

	if len(textNodes) == 0 {
		return nodes
	}
	firstNode := textNodes[0]
	if len(firstNode) > 0 {
		nodes = append(nodes, Plain(firstNode))
	}

	for i := 0; i < len(found); i++ {
		im := found[i]
		content := im[1]
		path := im[2]
		imageNode := Image{
			Content: SimpleParser(content),
			Path:    path,
		}
		textNode := textNodes[i+1]
		nodes = append(nodes, imageNode)
		if len(textNode) > 0 {
			nodes = append(nodes, Plain(textNode))
		}
	}

	return nodes
}

func HyperlinkParser(line string) []Node {
	// found := ImagePattern.FindAllString(line, -1)
	found := HyperlinkPattern.FindAllStringSubmatch(line, -1)
	textNodes := HyperlinkPattern.Split(line, -1)
	nodes := []Node{}

	if len(textNodes) == 0 {
		return nodes
	}
	firstNode := textNodes[0]
	if len(firstNode) > 0 {
		nodes = append(nodes, Plain(firstNode))
	}

	for i := 0; i < len(found); i++ {
		link := found[i]
		content := link[1]
		path := link[2]
		linkeNode := Hyperlink{
			Content: SimpleParser(content),
			Link:    path,
		}
		textNode := textNodes[i+1]
		nodes = append(nodes, linkeNode)
		if len(textNode) > 0 {
			nodes = append(nodes, Plain(textNode))
		}
	}

	return nodes
}

func nodePushFunc(nodes []Node, parseFunc func(string) []Node) []Node {
	newNodes := []Node{}
	for _, n := range nodes {
		switch v := n.(type) {
		case Plain:
			stringNode := string(v)
			newNodes = append(newNodes, parseFunc(stringNode)...)
		default:
			newNodes = append(newNodes, v)
		}
	}
	return newNodes
}

func NodeParser(nodes []Node) []Node {
	nodes = nodePushFunc(nodes, ImageParser)
	nodes = nodePushFunc(nodes, HyperlinkParser)
	nodes = nodePushFunc(nodes, SimpleParser)

	return nodes
}

func LineParser(line string) []Node {
	nodes := []Node{Plain(line)}

	return NodeParser(nodes)
}

func SetType(text string, TYPE int) Node {
	var typedText Node
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
