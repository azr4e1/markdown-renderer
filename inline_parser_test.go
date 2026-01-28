package markdownrenderer

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestImageParser(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Node
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []Node{},
		},
		{
			name:     "plain text only",
			input:    "hello world",
			expected: []Node{Plain("hello world")},
		},
		{
			name:     "single image",
			input:    "![alt text](image.png)",
			expected: []Node{Image{Content: []Node{Plain("alt text")}, Path: "image.png"}},
		},
		{
			name:     "image with plain text before",
			input:    "hello ![alt](img.png)",
			expected: []Node{Plain("hello "), Image{Content: []Node{Plain("alt")}, Path: "img.png"}},
		},
		{
			name:     "image with plain text after",
			input:    "![alt](img.png) world",
			expected: []Node{Image{Content: []Node{Plain("alt")}, Path: "img.png"}, Plain(" world")},
		},
		{
			name:     "image with text before and after",
			input:    "hello ![alt](img.png) world",
			expected: []Node{Plain("hello "), Image{Content: []Node{Plain("alt")}, Path: "img.png"}, Plain(" world")},
		},
		{
			name:     "multiple images",
			input:    "![a](1.png)![b](2.png)",
			expected: []Node{Image{Content: []Node{Plain("a")}, Path: "1.png"}, Image{Content: []Node{Plain("b")}, Path: "2.png"}},
		},
		{
			name:     "image with bold in alt text",
			input:    "![**bold**](img.png)",
			expected: []Node{Image{Content: []Node{Bold("bold")}, Path: "img.png"}},
		},
		{
			name:     "image with italic in alt text",
			input:    "![*italic*](img.png)",
			expected: []Node{Image{Content: []Node{Italic("italic")}, Path: "img.png"}},
		},
		{
			name:     "image with url path",
			input:    "![logo](https://example.com/logo.png)",
			expected: []Node{Image{Content: []Node{Plain("logo")}, Path: "https://example.com/logo.png"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ImageParser(tt.input)
			if diff := cmp.Diff(got, tt.expected); diff != "" {
				t.Errorf("ImageParser(%q)\n  got:      %v\n  expected: %v\n  Diff:      %s", tt.input, got, tt.expected, diff)
			}
		})
	}
}

func TestHyperlinkParser(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Node
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []Node{},
		},
		{
			name:     "plain text only",
			input:    "hello world",
			expected: []Node{Plain("hello world")},
		},
		{
			name:     "single link",
			input:    "[click here](https://example.com)",
			expected: []Node{Hyperlink{Content: []Node{Plain("click here")}, Link: "https://example.com"}},
		},
		{
			name:     "link with plain text before",
			input:    "visit [site](url.com)",
			expected: []Node{Plain("visit "), Hyperlink{Content: []Node{Plain("site")}, Link: "url.com"}},
		},
		{
			name:     "link with plain text after",
			input:    "[site](url.com) for info",
			expected: []Node{Hyperlink{Content: []Node{Plain("site")}, Link: "url.com"}, Plain(" for info")},
		},
		{
			name:     "link with text before and after",
			input:    "visit [site](url.com) now",
			expected: []Node{Plain("visit "), Hyperlink{Content: []Node{Plain("site")}, Link: "url.com"}, Plain(" now")},
		},
		{
			name:     "multiple links",
			input:    "[a](1.com)[b](2.com)",
			expected: []Node{Hyperlink{Content: []Node{Plain("a")}, Link: "1.com"}, Hyperlink{Content: []Node{Plain("b")}, Link: "2.com"}},
		},
		{
			name:     "link with bold in text",
			input:    "[**bold link**](url.com)",
			expected: []Node{Hyperlink{Content: []Node{Bold("bold link")}, Link: "url.com"}},
		},
		{
			name:     "link with italic in text",
			input:    "[*italic link*](url.com)",
			expected: []Node{Hyperlink{Content: []Node{Italic("italic link")}, Link: "url.com"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HyperlinkParser(tt.input)
			if diff := cmp.Diff(got, tt.expected); diff != "" {
				t.Errorf("HyperlinkParser(%q)\n  got:      %v\n  expected: %v\n  Diff:      %s", tt.input, got, tt.expected, diff)
			}
		})
	}
}

func TestSimpleParser(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Node
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []Node{Plain("")},
		},
		{
			name:     "single character",
			input:    "x",
			expected: []Node{Plain("x")},
		},
		{
			name:     "plain text only",
			input:    "hello world",
			expected: []Node{Plain("hello world")},
		},
		{
			name:     "bold with double asterisks",
			input:    "**bold**",
			expected: []Node{Bold("bold")},
		},
		{
			name:     "bold with double underscores",
			input:    "__bold__",
			expected: []Node{Bold("bold")},
		},
		{
			name:     "italic with single asterisk",
			input:    "*italic*",
			expected: []Node{Italic("italic")},
		},
		{
			name:     "italic with single underscore",
			input:    "_italic_",
			expected: []Node{Italic("italic")},
		},
		{
			name:     "underline with dash",
			input:    "-underline-",
			expected: []Node{Underline("underline")},
		},
		{
			name:     "inline code with backtick",
			input:    "`code`",
			expected: []Node{InlineCode("code")},
		},
		{
			name:     "crossed with tilde",
			input:    "~crossed~",
			expected: []Node{Crossed("crossed")},
		},
		{
			name:     "plain before bold",
			input:    "hello **world**",
			expected: []Node{Plain("hello "), Bold("world")},
		},
		{
			name:     "bold before plain",
			input:    "**hello** world",
			expected: []Node{Bold("hello"), Plain(" world")},
		},
		{
			name:     "original example with italic crossed and bold",
			input:    "*Ci ao c om*~ok ~**e va**",
			expected: []Node{Italic("Ci ao c om"), Crossed("ok "), Bold("e va")},
		},
		{
			name:     "italic then bold with plain between",
			input:    "*abc* **def**",
			expected: []Node{Italic("abc"), Plain(" "), Bold("def")},
		},
		{
			name:     "adjacent bold and italic no gap",
			input:    "**bold***italic*",
			expected: []Node{Bold("bold"), Italic("italic")},
		},
		{
			name:     "crossed then underline",
			input:    "~abc~ -def-",
			expected: []Node{Crossed("abc"), Plain(" "), Underline("def")},
		},
		{
			name:     "inline code then bold",
			input:    "`code` **bold**",
			expected: []Node{InlineCode("code"), Plain(" "), Bold("bold")},
		},
		{
			name:     "two italic segments",
			input:    "*a* *b*",
			expected: []Node{Italic("a"), Plain(" "), Italic("b")},
		},
		{
			name:     "two adjacent italic segments",
			input:    "*a**b*",
			expected: []Node{Italic("a"), Italic("b")},
		},
		{
			name:     "unclosed bold delimiter",
			input:    "**hello",
			expected: []Node{Bold("hello")},
		},
		{
			name:     "double underscore bold and underscore italic",
			input:    "__bold__ _italic_",
			expected: []Node{Bold("bold"), Plain(" "), Italic("italic")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SimpleParser(tt.input)
			if diff := cmp.Diff(got, tt.expected); diff != "" {
				t.Errorf("LineParser(%q)\n  got:      %v\n  expected: %v\n  Diff:      %s", tt.input, got, tt.expected, diff)
			}
		})
	}
}

func TestNodeParser(t *testing.T) {
	tests := []struct {
		name     string
		input    []Node
		expected []Node
	}{
		{
			name:     "empty slice",
			input:    []Node{},
			expected: []Node{},
		},
		{
			name:     "plain text only",
			input:    []Node{Plain("hello world")},
			expected: []Node{Plain("hello world")},
		},
		{
			name:     "bold text",
			input:    []Node{Plain("**bold**")},
			expected: []Node{Bold("bold")},
		},
		{
			name:     "single image",
			input:    []Node{Plain("![alt](img.png)")},
			expected: []Node{Image{Content: []Node{Plain("alt")}, Path: "img.png"}},
		},
		{
			name:     "single link",
			input:    []Node{Plain("[click](url.com)")},
			expected: []Node{Hyperlink{Content: []Node{Plain("click")}, Link: "url.com"}},
		},
		{
			name:     "image and link together",
			input:    []Node{Plain("![img](a.png) and [link](b.com)")},
			expected: []Node{Image{Content: []Node{Plain("img")}, Path: "a.png"}, Plain(" and "), Hyperlink{Content: []Node{Plain("link")}, Link: "b.com"}},
		},
		{
			name:     "link with bold text inside",
			input:    []Node{Plain("[**bold link**](url.com)")},
			expected: []Node{Hyperlink{Content: []Node{Bold("bold link")}, Link: "url.com"}},
		},
		{
			name:     "image with italic alt text",
			input:    []Node{Plain("![*italic*](img.png)")},
			expected: []Node{Image{Content: []Node{Italic("italic")}, Path: "img.png"}},
		},
		{
			name:     "mixed content with formatting",
			input:    []Node{Plain("hello **world** and [link](url.com)")},
			expected: []Node{Plain("hello "), Bold("world"), Plain(" and "), Hyperlink{Content: []Node{Plain("link")}, Link: "url.com"}},
		},
		{
			name:     "preserves non-plain nodes",
			input:    []Node{Bold("already bold"), Plain(" and **more**")},
			expected: []Node{Bold("already bold"), Plain(" and "), Bold("more")},
		},
		{
			name:     "complex mixed content",
			input:    []Node{Plain("![pic](a.png) *italic* [link](b.com) **bold**")},
			expected: []Node{Image{Content: []Node{Plain("pic")}, Path: "a.png"}, Plain(" "), Italic("italic"), Plain(" "), Hyperlink{Content: []Node{Plain("link")}, Link: "b.com"}, Plain(" "), Bold("bold")},
		},
		{
			name:     "multiple plain nodes",
			input:    []Node{Plain("**a**"), Plain(" "), Plain("**b**")},
			expected: []Node{Bold("a"), Plain(" "), Bold("b")},
		},
		{
			name:     "image then formatting in same line",
			input:    []Node{Plain("![alt](img.png) then `code`")},
			expected: []Node{Image{Content: []Node{Plain("alt")}, Path: "img.png"}, Plain(" then "), InlineCode("code")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NodeParser(tt.input)
			if diff := cmp.Diff(got, tt.expected); diff != "" {
				t.Errorf("LineParser(%v)\n  got:      %v\n  expected: %v\n  Diff:      %s", tt.input, got, tt.expected, diff)
			}
		})
	}
}
