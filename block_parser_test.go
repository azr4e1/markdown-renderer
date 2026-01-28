package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMarkdownToBlocks(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "single block",
			input:    "hello world",
			expected: []string{"hello world"},
		},
		{
			name:     "two blocks",
			input:    "block one\n\nblock two",
			expected: []string{"block one", "block two"},
		},
		{
			name:     "three blocks",
			input:    "one\n\ntwo\n\nthree",
			expected: []string{"one", "two", "three"},
		},
		{
			name:     "blocks with extra newlines",
			input:    "one\n\n\n\ntwo",
			expected: []string{"one", "two"},
		},
		{
			name:     "blocks with leading whitespace",
			input:    "  one  \n\n  two  ",
			expected: []string{"one", "two"},
		},
		{
			name:     "block with internal newline",
			input:    "line one\nline two\n\nblock two",
			expected: []string{"line one\nline two", "block two"},
		},
		{
			name:     "empty blocks filtered out",
			input:    "one\n\n\n\n\n\ntwo",
			expected: []string{"one", "two"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MarkdownToBlocks(tt.input)
			if diff := cmp.Diff(got, tt.expected); diff != "" {
				t.Errorf("MarkdownToBlocks(%q)\n  got:      %v\n  expected: %v\n  Diff:     %s", tt.input, got, tt.expected, diff)
			}
		})
	}
}

func TestIsHeader(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedLevel int
		expectedIsH   bool
	}{
		{
			name:          "h1 header",
			input:         "# Hello",
			expectedLevel: 1,
			expectedIsH:   true,
		},
		{
			name:          "h2 header",
			input:         "## Hello",
			expectedLevel: 2,
			expectedIsH:   true,
		},
		{
			name:          "h3 header",
			input:         "### Hello",
			expectedLevel: 3,
			expectedIsH:   true,
		},
		{
			name:          "h6 header",
			input:         "###### Hello",
			expectedLevel: 6,
			expectedIsH:   true,
		},
		{
			name:          "not a header plain text",
			input:         "Hello world",
			expectedLevel: 0,
			expectedIsH:   false,
		},
		{
			name:          "multiline is not header",
			input:         "# Hello\nworld",
			expectedLevel: 0,
			expectedIsH:   false,
		},
		{
			name:          "header without space",
			input:         "#Hello",
			expectedLevel: 1,
			expectedIsH:   false,
		},
		{
			name:          "header without space two levels",
			input:         "##Hello",
			expectedLevel: 2,
			expectedIsH:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			level, isH := isHeader(tt.input)
			if level != tt.expectedLevel || isH != tt.expectedIsH {
				t.Errorf("isHeader(%q) = (%d, %v), expected (%d, %v)", tt.input, level, isH, tt.expectedLevel, tt.expectedIsH)
			}
		})
	}
}

func TestIsBreak(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid break",
			input:    "---",
			expected: true,
		},
		{
			name:     "not a break plain text",
			input:    "hello",
			expected: false,
		},
		{
			name:     "not a break too few dashes",
			input:    "--",
			expected: false,
		},
		{
			name:     "not a break too many dashes",
			input:    "----",
			expected: false,
		},
		{
			name:     "not a break with spaces",
			input:    "- - -",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isBreak(tt.input)
			if got != tt.expected {
				t.Errorf("isBreak(%q) = %v, expected %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestIsCode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid code block",
			input:    "```\ncode here\n```",
			expected: true,
		},
		{
			name:     "code block with language",
			input:    "```go\nfunc main() {}\n```",
			expected: true,
		},
		{
			name:     "empty code block",
			input:    "``````",
			expected: true,
		},
		{
			name:     "not code block missing end",
			input:    "```\ncode here",
			expected: false,
		},
		{
			name:     "not code block missing start",
			input:    "code here\n```",
			expected: false,
		},
		{
			name:     "not code block plain text",
			input:    "hello world",
			expected: false,
		},
		{
			name:     "not code block too few backticks",
			input:    "``\nhello world\n``",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isCode(tt.input)
			if got != tt.expected {
				t.Errorf("isCode(%q) = %v, expected %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestIsQuote(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		expectedDelim   string
		expectedIsQuote bool
	}{
		{
			name:            "quote with greater than",
			input:           "> quoted text",
			expectedDelim:   "> ",
			expectedIsQuote: true,
		},
		{
			name:            "quote with two spaces",
			input:           "  indented text",
			expectedDelim:   "  ",
			expectedIsQuote: true,
		},
		{
			name:            "quote with tab",
			input:           "\tindented text",
			expectedDelim:   "\t",
			expectedIsQuote: true,
		},
		{
			name:            "multiline quote",
			input:           "> line one\n> line two",
			expectedDelim:   "> ",
			expectedIsQuote: true,
		},
		{
			name:            "not a quote plain text",
			input:           "hello world",
			expectedDelim:   "",
			expectedIsQuote: false,
		},
		{
			name:            "inconsistent prefix not a quote",
			input:           "> line one\nline two",
			expectedDelim:   "",
			expectedIsQuote: false,
		},
		{
			name:            "too short",
			input:           ">",
			expectedDelim:   "",
			expectedIsQuote: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delim, isQ := isQuote(tt.input)
			if delim != tt.expectedDelim || isQ != tt.expectedIsQuote {
				t.Errorf("isQuote(%q) = (%q, %v), expected (%q, %v)", tt.input, delim, isQ, tt.expectedDelim, tt.expectedIsQuote)
			}
		})
	}
}

func TestIsUnorderedList(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "single item with asterisk",
			input:    "* item one",
			expected: true,
		},
		{
			name:     "single item with dash",
			input:    "- item one",
			expected: true,
		},
		{
			name:     "multiple items with asterisk",
			input:    "* item one\n* item two",
			expected: true,
		},
		{
			name:     "multiple items with dash",
			input:    "- item one\n- item two",
			expected: true,
		},
		{
			name:     "mixed asterisk and dash",
			input:    "* item one\n- item two",
			expected: true,
		},
		{
			name:     "not a list plain text",
			input:    "hello world",
			expected: false,
		},
		{
			name:     "not a list missing space after asterisk",
			input:    "*item",
			expected: false,
		},
		{
			name:     "not a list missing bullet point",
			input:    "* item\nitem2",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isUnorderedList(tt.input)
			if got != tt.expected {
				t.Errorf("isUnorderedList(%q) = %v, expected %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestIsOrderedList(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "single item",
			input:    "1. item one",
			expected: true,
		},
		{
			name:     "multiple items sequential",
			input:    "1. item one\n2. item two\n3. item three",
			expected: true,
		},
		{
			name:     "all ones prefix",
			input:    "1. item one\n1. item two",
			expected: true,
		},
		{
			name:     "not a list plain text",
			input:    "hello world",
			expected: false,
		},
		{
			name:     "not a list wrong number",
			input:    "1. item one\n3. item three",
			expected: false,
		},
		{
			name:     "not a list missing space",
			input:    "1.item",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isOrderedList(tt.input)
			if got != tt.expected {
				t.Errorf("isOrderedList(%q) = %v, expected %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestBlockParser(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Node
	}{
		{
			name:     "header h1",
			input:    "# Hello",
			expected: Header{Content: []Node{Plain("Hello")}, Level: 1},
		},
		{
			name:     "header h2",
			input:    "## World",
			expected: Header{Content: []Node{Plain("World")}, Level: 2},
		},
		{
			name:     "break",
			input:    "---",
			expected: Break(true),
		},
		{
			name:     "code block",
			input:    "```\nfmt.Println()\n```",
			expected: Code("\nfmt.Println()\n"),
		},
		{
			name:     "quote",
			input:    "> quoted text",
			expected: Quote([]Node{Plain("quoted text")}),
		},
		{
			name:  "unordered list",
			input: "* item one\n* item two",
			expected: UnorderedList{
				UnorderedItem([]Node{Plain("item one")}),
				UnorderedItem([]Node{Plain("item two")}),
			},
		},
		{
			name:  "ordered list",
			input: "1. first\n2. second",
			expected: OrderedList{
				OrderedItem([]Node{Plain("first")}),
				OrderedItem([]Node{Plain("second")}),
			},
		},
		{
			name:     "paragraph",
			input:    "Just some text",
			expected: Paragraph([]Node{Plain("Just some text")}),
		},
		{
			name:     "paragraph with bold",
			input:    "Some **bold** text",
			expected: Paragraph([]Node{Plain("Some "), Bold("bold"), Plain(" text")}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BlockParser(tt.input)
			if diff := cmp.Diff(got, tt.expected); diff != "" {
				t.Errorf("BlockParser(%q)\n  got:      %v\n  expected: %v\n  Diff:     %s", tt.input, got, tt.expected, diff)
			}
		})
	}
}

func TestCodeify(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Code
	}{
		{
			name:     "simple code",
			input:    "```\ncode\n```",
			expected: Code("\ncode\n"),
		},
		{
			name:     "code with language",
			input:    "```go\nfunc main() {}\n```",
			expected: Code("go\nfunc main() {}\n"),
		},
		{
			name:     "empty code",
			input:    "``````",
			expected: Code(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := codeify(tt.input)
			if got != tt.expected {
				t.Errorf("codeify(%q) = %v, expected %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestQuoteify(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		delimiter string
		expected  Quote
	}{
		{
			name:      "single line quote",
			input:     "> hello",
			delimiter: "> ",
			expected:  Quote([]Node{Plain("hello")}),
		},
		{
			name:      "multiline quote",
			input:     "> line one\n> line two",
			delimiter: "> ",
			expected:  Quote([]Node{Plain("line one\nline two")}),
		},
		{
			name:      "quote with bold",
			input:     "> **bold** text",
			delimiter: "> ",
			expected:  Quote([]Node{Bold("bold"), Plain(" text")}),
		},
		{
			name:      "quote with italic",
			input:     "> *italic* here",
			delimiter: "> ",
			expected:  Quote([]Node{Italic("italic"), Plain(" here")}),
		},
		{
			name:      "tab indented quote",
			input:     "\tindented line",
			delimiter: "\t",
			expected:  Quote([]Node{Plain("indented line")}),
		},
		{
			name:      "space indented quote",
			input:     "  indented line",
			delimiter: "  ",
			expected:  Quote([]Node{Plain("indented line")}),
		},
		{
			name:      "multiline tab quote",
			input:     "\tline one\n\tline two",
			delimiter: "\t",
			expected:  Quote([]Node{Plain("line one\nline two")}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := quoteify(tt.input, tt.delimiter)
			if diff := cmp.Diff(got, tt.expected); diff != "" {
				t.Errorf("quoteify(%q, %q)\n  got:      %v\n  expected: %v\n  Diff:     %s", tt.input, tt.delimiter, got, tt.expected, diff)
			}
		})
	}
}

func TestUlistify(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected UnorderedList
	}{
		{
			name:  "single item with asterisk",
			input: "* item one",
			expected: UnorderedList{
				UnorderedItem([]Node{Plain("item one")}),
			},
		},
		{
			name:  "single item with dash",
			input: "- item one",
			expected: UnorderedList{
				UnorderedItem([]Node{Plain("item one")}),
			},
		},
		{
			name:  "multiple items with asterisk",
			input: "* first\n* second\n* third",
			expected: UnorderedList{
				UnorderedItem([]Node{Plain("first")}),
				UnorderedItem([]Node{Plain("second")}),
				UnorderedItem([]Node{Plain("third")}),
			},
		},
		{
			name:  "multiple items with dash",
			input: "- first\n- second",
			expected: UnorderedList{
				UnorderedItem([]Node{Plain("first")}),
				UnorderedItem([]Node{Plain("second")}),
			},
		},
		{
			name:  "mixed asterisk and dash",
			input: "* first\n- second",
			expected: UnorderedList{
				UnorderedItem([]Node{Plain("first")}),
				UnorderedItem([]Node{Plain("second")}),
			},
		},
		{
			name:  "item with bold",
			input: "* **bold** item",
			expected: UnorderedList{
				UnorderedItem([]Node{Bold("bold"), Plain(" item")}),
			},
		},
		{
			name:  "item with italic",
			input: "* *italic* item",
			expected: UnorderedList{
				UnorderedItem([]Node{Italic("italic"), Plain(" item")}),
			},
		},
		{
			name:  "item with link",
			input: "* [link](url.com)",
			expected: UnorderedList{
				UnorderedItem([]Node{Hyperlink{Content: []Node{Plain("link")}, Link: "url.com"}}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ulistify(tt.input)
			if diff := cmp.Diff(got, tt.expected); diff != "" {
				t.Errorf("ulistify(%q)\n  got:      %v\n  expected: %v\n  Diff:     %s", tt.input, got, tt.expected, diff)
			}
		})
	}
}

func TestOlistify(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected OrderedList
	}{
		{
			name:  "single item",
			input: "1. item one",
			expected: OrderedList{
				OrderedItem([]Node{Plain("item one")}),
			},
		},
		{
			name:  "multiple items sequential",
			input: "1. first\n2. second\n3. third",
			expected: OrderedList{
				OrderedItem([]Node{Plain("first")}),
				OrderedItem([]Node{Plain("second")}),
				OrderedItem([]Node{Plain("third")}),
			},
		},
		{
			name:  "all ones prefix",
			input: "1. first\n1. second",
			expected: OrderedList{
				OrderedItem([]Node{Plain("first")}),
				OrderedItem([]Node{Plain("second")}),
			},
		},
		{
			name:  "item with bold",
			input: "1. **bold** item",
			expected: OrderedList{
				OrderedItem([]Node{Bold("bold"), Plain(" item")}),
			},
		},
		{
			name:  "item with italic",
			input: "1. *italic* item",
			expected: OrderedList{
				OrderedItem([]Node{Italic("italic"), Plain(" item")}),
			},
		},
		{
			name:  "item with link",
			input: "1. [link](url.com)",
			expected: OrderedList{
				OrderedItem([]Node{Hyperlink{Content: []Node{Plain("link")}, Link: "url.com"}}),
			},
		},
		{
			name:  "five items",
			input: "1. one\n2. two\n3. three\n4. four\n5. five",
			expected: OrderedList{
				OrderedItem([]Node{Plain("one")}),
				OrderedItem([]Node{Plain("two")}),
				OrderedItem([]Node{Plain("three")}),
				OrderedItem([]Node{Plain("four")}),
				OrderedItem([]Node{Plain("five")}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := olistify(tt.input)
			if diff := cmp.Diff(got, tt.expected); diff != "" {
				t.Errorf("olistify(%q)\n  got:      %v\n  expected: %v\n  Diff:     %s", tt.input, got, tt.expected, diff)
			}
		})
	}
}
