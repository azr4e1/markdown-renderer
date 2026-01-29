package markdownrenderer

import "testing"

func TestHTMLPlainRender(t *testing.T) {
	tests := []struct {
		name     string
		input    HTMLPlain
		expected string
	}{
		{name: "empty", input: HTMLPlain(""), expected: ""},
		{name: "text", input: HTMLPlain("hello"), expected: "hello"},
		{name: "special chars", input: HTMLPlain("a & b"), expected: "a & b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.HTMLRender()
			if got != tt.expected {
				t.Errorf("HTMLPlain(%q).HTMLRender() = %q, expected %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestHTMLBoldRender(t *testing.T) {
	tests := []struct {
		name     string
		input    HTMLBold
		expected string
	}{
		{name: "simple", input: HTMLBold("bold"), expected: "<b>bold</b>"},
		{name: "empty", input: HTMLBold(""), expected: "<b></b>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.HTMLRender()
			if got != tt.expected {
				t.Errorf("HTMLBold(%q).HTMLRender() = %q, expected %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestHTMLItalicRender(t *testing.T) {
	tests := []struct {
		name     string
		input    HTMLItalic
		expected string
	}{
		{name: "simple", input: HTMLItalic("italic"), expected: "<i>italic</i>"},
		{name: "empty", input: HTMLItalic(""), expected: "<i></i>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.HTMLRender()
			if got != tt.expected {
				t.Errorf("HTMLItalic(%q).HTMLRender() = %q, expected %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestHTMLUnderlineRender(t *testing.T) {
	tests := []struct {
		name     string
		input    HTMLUnderline
		expected string
	}{
		{name: "simple", input: HTMLUnderline("underline"), expected: "<u>underline</u>"},
		{name: "empty", input: HTMLUnderline(""), expected: "<u></u>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.HTMLRender()
			if got != tt.expected {
				t.Errorf("HTMLUnderline(%q).HTMLRender() = %q, expected %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestHTMLInlineCodeRender(t *testing.T) {
	tests := []struct {
		name     string
		input    HTMLInlineCode
		expected string
	}{
		{name: "simple", input: HTMLInlineCode("code"), expected: "<code>code</code>"},
		{name: "empty", input: HTMLInlineCode(""), expected: "<code></code>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.HTMLRender()
			if got != tt.expected {
				t.Errorf("HTMLInlineCode(%q).HTMLRender() = %q, expected %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestHTMLCrossedRender(t *testing.T) {
	tests := []struct {
		name     string
		input    HTMLCrossed
		expected string
	}{
		{name: "simple", input: HTMLCrossed("crossed"), expected: "<strike>crossed</strike>"},
		{name: "empty", input: HTMLCrossed(""), expected: "<strike></strike>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.HTMLRender()
			if got != tt.expected {
				t.Errorf("HTMLCrossed(%q).HTMLRender() = %q, expected %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestHTMLHyperlinkRender(t *testing.T) {
	tests := []struct {
		name     string
		input    HTMLHyperlink
		expected string
	}{
		{
			name:     "plain content",
			input:    HTMLHyperlink{Content: []HTMLNode{HTMLPlain("click")}, Link: "https://example.com"},
			expected: "<a href='https://example.com'>click</a>",
		},
		{
			name:     "bold content",
			input:    HTMLHyperlink{Content: []HTMLNode{HTMLBold("bold")}, Link: "url.com"},
			expected: "<a href='url.com'><b>bold</b></a>",
		},
		{
			name:     "empty content",
			input:    HTMLHyperlink{Content: []HTMLNode{}, Link: "url.com"},
			expected: "<a href='url.com'></a>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.HTMLRender()
			if got != tt.expected {
				t.Errorf("HTMLHyperlink.HTMLRender() = %q, expected %q", got, tt.expected)
			}
		})
	}
}

func TestHTMLImageRender(t *testing.T) {
	tests := []struct {
		name     string
		input    HTMLImage
		expected string
	}{
		{
			name:     "plain alt text",
			input:    HTMLImage{Content: []HTMLNode{HTMLPlain("alt")}, Path: "img.png"},
			expected: "<img src='img.png'>alt</img>",
		},
		{
			name:     "empty alt text",
			input:    HTMLImage{Content: []HTMLNode{}, Path: "img.png"},
			expected: "<img src='img.png'></img>",
		},
		{
			name:     "url path",
			input:    HTMLImage{Content: []HTMLNode{HTMLPlain("logo")}, Path: "https://example.com/logo.png"},
			expected: "<img src='https://example.com/logo.png'>logo</img>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.HTMLRender()
			if got != tt.expected {
				t.Errorf("HTMLImage.HTMLRender() = %q, expected %q", got, tt.expected)
			}
		})
	}
}

func TestHTMLDivRender(t *testing.T) {
	tests := []struct {
		name     string
		input    HTMLDiv
		expected string
	}{
		{
			name:     "single paragraph",
			input:    HTMLDiv{HTMLParagraph{HTMLPlain("hello")}},
			expected: "<div><p>hello</p></div>",
		},
		{
			name:     "empty div",
			input:    HTMLDiv{},
			expected: "<div></div>",
		},
		{
			name:     "multiple children",
			input:    HTMLDiv{HTMLParagraph{HTMLPlain("one")}, HTMLParagraph{HTMLPlain("two")}},
			expected: "<div><p>one</p><p>two</p></div>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.HTMLRender()
			if got != tt.expected {
				t.Errorf("HTMLDiv.HTMLRender() = %q, expected %q", got, tt.expected)
			}
		})
	}
}

func TestHTMLHeaderRender(t *testing.T) {
	tests := []struct {
		name     string
		input    HTMLHeader
		expected string
	}{
		{
			name:     "h1",
			input:    HTMLHeader{Content: []HTMLNode{HTMLPlain("Title")}, Level: 1},
			expected: "<h1>Title</h1>",
		},
		{
			name:     "h3",
			input:    HTMLHeader{Content: []HTMLNode{HTMLPlain("Subtitle")}, Level: 3},
			expected: "<h3>Subtitle</h3>",
		},
		{
			name:     "h2 with bold",
			input:    HTMLHeader{Content: []HTMLNode{HTMLBold("Bold Title")}, Level: 2},
			expected: "<h2><b>Bold Title</b></h2>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.HTMLRender()
			if got != tt.expected {
				t.Errorf("HTMLHeader.HTMLRender() = %q, expected %q", got, tt.expected)
			}
		})
	}
}

func TestHTMLParagraphRender(t *testing.T) {
	tests := []struct {
		name     string
		input    HTMLParagraph
		expected string
	}{
		{
			name:     "plain text",
			input:    HTMLParagraph{HTMLPlain("hello world")},
			expected: "<p>hello world</p>",
		},
		{
			name:     "mixed content",
			input:    HTMLParagraph{HTMLPlain("hello "), HTMLBold("bold"), HTMLPlain(" world")},
			expected: "<p>hello <b>bold</b> world</p>",
		},
		{
			name:     "empty",
			input:    HTMLParagraph{},
			expected: "<p></p>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.HTMLRender()
			if got != tt.expected {
				t.Errorf("HTMLParagraph.HTMLRender() = %q, expected %q", got, tt.expected)
			}
		})
	}
}

func TestHTMLCodeRender(t *testing.T) {
	tests := []struct {
		name     string
		input    HTMLCode
		expected string
	}{
		{
			name:     "simple code",
			input:    HTMLCode("fmt.Println()"),
			expected: "<pre><code>fmt.Println()</code></pre>",
		},
		{
			name:     "empty",
			input:    HTMLCode(""),
			expected: "<pre><code></code></pre>",
		},
		{
			name:     "multiline",
			input:    HTMLCode("line1\nline2"),
			expected: "<pre><code>line1\nline2</code></pre>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.HTMLRender()
			if got != tt.expected {
				t.Errorf("HTMLCode(%q).HTMLRender() = %q, expected %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestHTMLQuoteRender(t *testing.T) {
	tests := []struct {
		name     string
		input    HTMLQuote
		expected string
	}{
		{
			name:     "plain text",
			input:    HTMLQuote{HTMLPlain("quoted")},
			expected: "<blockquote>quoted</blockquote>",
		},
		{
			name:     "with bold",
			input:    HTMLQuote{HTMLPlain("hello "), HTMLBold("bold")},
			expected: "<blockquote>hello <b>bold</b></blockquote>",
		},
		{
			name:     "empty",
			input:    HTMLQuote{},
			expected: "<blockquote></blockquote>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.HTMLRender()
			if got != tt.expected {
				t.Errorf("HTMLQuote.HTMLRender() = %q, expected %q", got, tt.expected)
			}
		})
	}
}

func TestHTMLBreakRender(t *testing.T) {
	got := HTMLBreak(true).HTMLRender()
	expected := "<br>"
	if got != expected {
		t.Errorf("HTMLBreak.HTMLRender() = %q, expected %q", got, expected)
	}
}

func TestHTMLOrderedListRender(t *testing.T) {
	tests := []struct {
		name     string
		input    HTMLOrderedList
		expected string
	}{
		{
			name: "single item",
			input: HTMLOrderedList{
				HTMLOrderedItem{HTMLPlain("first")},
			},
			expected: "<ol><li>first</li></ol>",
		},
		{
			name: "multiple items",
			input: HTMLOrderedList{
				HTMLOrderedItem{HTMLPlain("first")},
				HTMLOrderedItem{HTMLPlain("second")},
				HTMLOrderedItem{HTMLPlain("third")},
			},
			expected: "<ol><li>first</li><li>second</li><li>third</li></ol>",
		},
		{
			name:     "empty list",
			input:    HTMLOrderedList{},
			expected: "<ol></ol>",
		},
		{
			name: "item with bold",
			input: HTMLOrderedList{
				HTMLOrderedItem{HTMLBold("bold"), HTMLPlain(" item")},
			},
			expected: "<ol><li><b>bold</b> item</li></ol>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.HTMLRender()
			if got != tt.expected {
				t.Errorf("HTMLOrderedList.HTMLRender() = %q, expected %q", got, tt.expected)
			}
		})
	}
}

func TestHTMLUnorderedListRender(t *testing.T) {
	tests := []struct {
		name     string
		input    HTMLUnorderedList
		expected string
	}{
		{
			name: "single item",
			input: HTMLUnorderedList{
				HTMLUnorderedItem{HTMLPlain("first")},
			},
			expected: "<ul><li>first</li></ul>",
		},
		{
			name: "multiple items",
			input: HTMLUnorderedList{
				HTMLUnorderedItem{HTMLPlain("first")},
				HTMLUnorderedItem{HTMLPlain("second")},
			},
			expected: "<ul><li>first</li><li>second</li></ul>",
		},
		{
			name:     "empty list",
			input:    HTMLUnorderedList{},
			expected: "<ul></ul>",
		},
		{
			name: "item with italic",
			input: HTMLUnorderedList{
				HTMLUnorderedItem{HTMLItalic("italic"), HTMLPlain(" item")},
			},
			expected: "<ul><li><i>italic</i> item</li></ul>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.HTMLRender()
			if got != tt.expected {
				t.Errorf("HTMLUnorderedList.HTMLRender() = %q, expected %q", got, tt.expected)
			}
		})
	}
}

func TestHtmlRender(t *testing.T) {
	tests := []struct {
		name     string
		input    []HTMLNode
		expected string
	}{
		{
			name:     "empty",
			input:    []HTMLNode{},
			expected: "",
		},
		{
			name:     "single plain",
			input:    []HTMLNode{HTMLPlain("hello")},
			expected: "hello",
		},
		{
			name:     "mixed nodes",
			input:    []HTMLNode{HTMLPlain("hello "), HTMLBold("world")},
			expected: "hello <b>world</b>",
		},
		{
			name:     "multiple inline types",
			input:    []HTMLNode{HTMLItalic("a"), HTMLPlain(" "), HTMLCrossed("b"), HTMLPlain(" "), HTMLInlineCode("c")},
			expected: "<i>a</i> <strike>b</strike> <code>c</code>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := htmlRender(tt.input)
			if got != tt.expected {
				t.Errorf("htmlRender() = %q, expected %q", got, tt.expected)
			}
		})
	}
}
