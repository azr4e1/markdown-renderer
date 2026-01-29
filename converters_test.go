package markdownrenderer

import "testing"

func TestMarkdownToHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "<div></div>",
		},
		{
			name:     "plain paragraph",
			input:    "hello world",
			expected: "<div><p>hello world</p></div>",
		},
		{
			name:     "h1 header",
			input:    "# Title",
			expected: "<div><h1>Title</h1></div>",
		},
		{
			name:     "h2 header",
			input:    "## Subtitle",
			expected: "<div><h2>Subtitle</h2></div>",
		},
		{
			name:     "header with bold",
			input:    "# **Bold** Title",
			expected: "<div><h1><b>Bold</b> Title</h1></div>",
		},
		{
			name:     "break",
			input:    "---",
			expected: "<div><br></div>",
		},
		{
			name:     "code block",
			input:    "```\nfmt.Println()\n```",
			expected: "<div><pre><code>\nfmt.Println()\n</code></pre></div>",
		},
		{
			name:     "quote",
			input:    "> quoted text",
			expected: "<div><blockquote>quoted text</blockquote></div>",
		},
		{
			name:     "multiline quote",
			input:    "> line one\n> line two",
			expected: "<div><blockquote>line one\nline two</blockquote></div>",
		},
		{
			name:     "unordered list",
			input:    "* item one\n* item two",
			expected: "<div><ul><li>item one</li><li>item two</li></ul></div>",
		},
		{
			name:     "ordered list",
			input:    "1. first\n2. second",
			expected: "<div><ol><li>first</li><li>second</li></ol></div>",
		},
		{
			name:     "bold text",
			input:    "**bold**",
			expected: "<div><p><b>bold</b></p></div>",
		},
		{
			name:     "italic text",
			input:    "*italic*",
			expected: "<div><p><i>italic</i></p></div>",
		},
		{
			name:     "inline code",
			input:    "`code`",
			expected: "<div><p><code>code</code></p></div>",
		},
		{
			name:     "crossed text",
			input:    "~crossed~",
			expected: "<div><p><strike>crossed</strike></p></div>",
		},
		{
			name:     "underline text",
			input:    "-underline-",
			expected: "<div><p><u>underline</u></p></div>",
		},
		{
			name:     "hyperlink",
			input:    "[click](https://example.com)",
			expected: "<div><p><a href='https://example.com'>click</a></p></div>",
		},
		{
			name:     "image",
			input:    "![alt](img.png)",
			expected: "<div><p><img src='img.png'>alt</img></p></div>",
		},
		{
			name:     "paragraph with mixed inline",
			input:    "hello **bold** and *italic* text",
			expected: "<div><p>hello <b>bold</b> and <i>italic</i> text</p></div>",
		},
		{
			name:     "two paragraphs",
			input:    "first paragraph\n\nsecond paragraph",
			expected: "<div><p>first paragraph</p><p>second paragraph</p></div>",
		},
		{
			name:     "header then paragraph",
			input:    "# Title\n\nSome text here",
			expected: "<div><h1>Title</h1><p>Some text here</p></div>",
		},
		{
			name:     "paragraph then break then paragraph",
			input:    "before\n\n---\n\nafter",
			expected: "<div><p>before</p><br><p>after</p></div>",
		},
		{
			name:     "header then list",
			input:    "# Shopping\n\n* apples\n* bananas",
			expected: "<div><h1>Shopping</h1><ul><li>apples</li><li>bananas</li></ul></div>",
		},
		{
			name:     "paragraph then code block",
			input:    "Example:\n\n```\nx := 1\n```",
			expected: "<div><p>Example:</p><pre><code>\nx := 1\n</code></pre></div>",
		},
		{
			name:     "link with bold text",
			input:    "[**bold link**](url.com)",
			expected: "<div><p><a href='url.com'><b>bold link</b></a></p></div>",
		},
		{
			name:     "ordered list with formatting",
			input:    "1. **bold** item\n2. *italic* item",
			expected: "<div><ol><li><b>bold</b> item</li><li><i>italic</i> item</li></ol></div>",
		},
		{
			name:     "quote with formatting",
			input:    "> **bold** and *italic*",
			expected: "<div><blockquote><b>bold</b> and <i>italic</i></blockquote></div>",
		},
		{
			name:     "three blocks mixed",
			input:    "# Header\n\nParagraph text\n\n* list item",
			expected: "<div><h1>Header</h1><p>Paragraph text</p><ul><li>list item</li></ul></div>",
		},
		{
			name:     "image and link in paragraph",
			input:    "See ![pic](a.png) and [link](b.com)",
			expected: "<div><p>See <img src='a.png'>pic</img> and <a href='b.com'>link</a></p></div>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MarkdownToHTML(tt.input)
			rendered := got.HTMLRender()
			if rendered != tt.expected {
				t.Errorf("MarkdownToHTML(%q).HTMLRender()\n  got:      %q\n  expected: %q", tt.input, rendered, tt.expected)
			}
		})
	}
}
