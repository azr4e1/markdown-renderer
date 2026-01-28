package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestImageParser(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Text
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []Text{},
		},
		{
			name:     "plain text only",
			input:    "hello world",
			expected: []Text{Plain("hello world")},
		},
		{
			name:     "single image",
			input:    "![alt text](image.png)",
			expected: []Text{Image{Content: []Text{Plain("alt text")}, Path: "image.png"}},
		},
		{
			name:     "image with plain text before",
			input:    "hello ![alt](img.png)",
			expected: []Text{Plain("hello "), Image{Content: []Text{Plain("alt")}, Path: "img.png"}},
		},
		{
			name:     "image with plain text after",
			input:    "![alt](img.png) world",
			expected: []Text{Image{Content: []Text{Plain("alt")}, Path: "img.png"}, Plain(" world")},
		},
		{
			name:     "image with text before and after",
			input:    "hello ![alt](img.png) world",
			expected: []Text{Plain("hello "), Image{Content: []Text{Plain("alt")}, Path: "img.png"}, Plain(" world")},
		},
		{
			name:     "multiple images",
			input:    "![a](1.png)![b](2.png)",
			expected: []Text{Image{Content: []Text{Plain("a")}, Path: "1.png"}, Image{Content: []Text{Plain("b")}, Path: "2.png"}},
		},
		{
			name:     "image with bold in alt text",
			input:    "![**bold**](img.png)",
			expected: []Text{Image{Content: []Text{Bold("bold")}, Path: "img.png"}},
		},
		{
			name:     "image with italic in alt text",
			input:    "![*italic*](img.png)",
			expected: []Text{Image{Content: []Text{Italic("italic")}, Path: "img.png"}},
		},
		{
			name:     "image with url path",
			input:    "![logo](https://example.com/logo.png)",
			expected: []Text{Image{Content: []Text{Plain("logo")}, Path: "https://example.com/logo.png"}},
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
		expected []Text
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []Text{},
		},
		{
			name:     "plain text only",
			input:    "hello world",
			expected: []Text{Plain("hello world")},
		},
		{
			name:     "single link",
			input:    "[click here](https://example.com)",
			expected: []Text{Hyperlink{Content: []Text{Plain("click here")}, Link: "https://example.com"}},
		},
		{
			name:     "link with plain text before",
			input:    "visit [site](url.com)",
			expected: []Text{Plain("visit "), Hyperlink{Content: []Text{Plain("site")}, Link: "url.com"}},
		},
		{
			name:     "link with plain text after",
			input:    "[site](url.com) for info",
			expected: []Text{Hyperlink{Content: []Text{Plain("site")}, Link: "url.com"}, Plain(" for info")},
		},
		{
			name:     "link with text before and after",
			input:    "visit [site](url.com) now",
			expected: []Text{Plain("visit "), Hyperlink{Content: []Text{Plain("site")}, Link: "url.com"}, Plain(" now")},
		},
		{
			name:     "multiple links",
			input:    "[a](1.com)[b](2.com)",
			expected: []Text{Hyperlink{Content: []Text{Plain("a")}, Link: "1.com"}, Hyperlink{Content: []Text{Plain("b")}, Link: "2.com"}},
		},
		{
			name:     "link with bold in text",
			input:    "[**bold link**](url.com)",
			expected: []Text{Hyperlink{Content: []Text{Bold("bold link")}, Link: "url.com"}},
		},
		{
			name:     "link with italic in text",
			input:    "[*italic link*](url.com)",
			expected: []Text{Hyperlink{Content: []Text{Italic("italic link")}, Link: "url.com"}},
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
		expected []Text
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []Text{Plain("")},
		},
		{
			name:     "single character",
			input:    "x",
			expected: []Text{Plain("x")},
		},
		{
			name:     "plain text only",
			input:    "hello world",
			expected: []Text{Plain("hello world")},
		},
		{
			name:     "bold with double asterisks",
			input:    "**bold**",
			expected: []Text{Bold("bold")},
		},
		{
			name:     "bold with double underscores",
			input:    "__bold__",
			expected: []Text{Bold("bold")},
		},
		{
			name:     "italic with single asterisk",
			input:    "*italic*",
			expected: []Text{Italic("italic")},
		},
		{
			name:     "italic with single underscore",
			input:    "_italic_",
			expected: []Text{Italic("italic")},
		},
		{
			name:     "underline with dash",
			input:    "-underline-",
			expected: []Text{Underline("underline")},
		},
		{
			name:     "inline code with backtick",
			input:    "`code`",
			expected: []Text{InlineCode("code")},
		},
		{
			name:     "crossed with tilde",
			input:    "~crossed~",
			expected: []Text{Crossed("crossed")},
		},
		{
			name:     "plain before bold",
			input:    "hello **world**",
			expected: []Text{Plain("hello "), Bold("world")},
		},
		{
			name:     "bold before plain",
			input:    "**hello** world",
			expected: []Text{Bold("hello"), Plain(" world")},
		},
		{
			name:     "original example with italic crossed and bold",
			input:    "*Ci ao c om*~ok ~**e va**",
			expected: []Text{Italic("Ci ao c om"), Crossed("ok "), Bold("e va")},
		},
		{
			name:     "italic then bold with plain between",
			input:    "*abc* **def**",
			expected: []Text{Italic("abc"), Plain(" "), Bold("def")},
		},
		{
			name:     "adjacent bold and italic no gap",
			input:    "**bold***italic*",
			expected: []Text{Bold("bold"), Italic("italic")},
		},
		{
			name:     "crossed then underline",
			input:    "~abc~ -def-",
			expected: []Text{Crossed("abc"), Plain(" "), Underline("def")},
		},
		{
			name:     "inline code then bold",
			input:    "`code` **bold**",
			expected: []Text{InlineCode("code"), Plain(" "), Bold("bold")},
		},
		{
			name:     "two italic segments",
			input:    "*a* *b*",
			expected: []Text{Italic("a"), Plain(" "), Italic("b")},
		},
		{
			name:     "two adjacent italic segments",
			input:    "*a**b*",
			expected: []Text{Italic("a"), Italic("b")},
		},
		{
			name:     "unclosed bold delimiter",
			input:    "**hello",
			expected: []Text{Bold("hello")},
		},
		{
			name:     "double underscore bold and underscore italic",
			input:    "__bold__ _italic_",
			expected: []Text{Bold("bold"), Plain(" "), Italic("italic")},
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
		input    []Text
		expected []Text
	}{
		{
			name:     "empty slice",
			input:    []Text{},
			expected: []Text{},
		},
		{
			name:     "plain text only",
			input:    []Text{Plain("hello world")},
			expected: []Text{Plain("hello world")},
		},
		{
			name:     "bold text",
			input:    []Text{Plain("**bold**")},
			expected: []Text{Bold("bold")},
		},
		{
			name:     "single image",
			input:    []Text{Plain("![alt](img.png)")},
			expected: []Text{Image{Content: []Text{Plain("alt")}, Path: "img.png"}},
		},
		{
			name:     "single link",
			input:    []Text{Plain("[click](url.com)")},
			expected: []Text{Hyperlink{Content: []Text{Plain("click")}, Link: "url.com"}},
		},
		{
			name:     "image and link together",
			input:    []Text{Plain("![img](a.png) and [link](b.com)")},
			expected: []Text{Image{Content: []Text{Plain("img")}, Path: "a.png"}, Plain(" and "), Hyperlink{Content: []Text{Plain("link")}, Link: "b.com"}},
		},
		{
			name:     "link with bold text inside",
			input:    []Text{Plain("[**bold link**](url.com)")},
			expected: []Text{Hyperlink{Content: []Text{Bold("bold link")}, Link: "url.com"}},
		},
		{
			name:     "image with italic alt text",
			input:    []Text{Plain("![*italic*](img.png)")},
			expected: []Text{Image{Content: []Text{Italic("italic")}, Path: "img.png"}},
		},
		{
			name:     "mixed content with formatting",
			input:    []Text{Plain("hello **world** and [link](url.com)")},
			expected: []Text{Plain("hello "), Bold("world"), Plain(" and "), Hyperlink{Content: []Text{Plain("link")}, Link: "url.com"}},
		},
		{
			name:     "preserves non-plain nodes",
			input:    []Text{Bold("already bold"), Plain(" and **more**")},
			expected: []Text{Bold("already bold"), Plain(" and "), Bold("more")},
		},
		{
			name:     "complex mixed content",
			input:    []Text{Plain("![pic](a.png) *italic* [link](b.com) **bold**")},
			expected: []Text{Image{Content: []Text{Plain("pic")}, Path: "a.png"}, Plain(" "), Italic("italic"), Plain(" "), Hyperlink{Content: []Text{Plain("link")}, Link: "b.com"}, Plain(" "), Bold("bold")},
		},
		{
			name:     "multiple plain nodes",
			input:    []Text{Plain("**a**"), Plain(" "), Plain("**b**")},
			expected: []Text{Bold("a"), Plain(" "), Bold("b")},
		},
		{
			name:     "image then formatting in same line",
			input:    []Text{Plain("![alt](img.png) then `code`")},
			expected: []Text{Image{Content: []Text{Plain("alt")}, Path: "img.png"}, Plain(" then "), InlineCode("code")},
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
