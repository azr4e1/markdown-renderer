package markdownrenderer

import "fmt"

type HTMLNode interface {
	HTMLRender() string
}

func htmlRender(nodes []HTMLNode) string {
	render := ""
	for _, h := range nodes {
		render += h.HTMLRender()
	}

	return render
}

// Leaves

type HTMLPlain string
type HTMLBold string
type HTMLItalic string
type HTMLUnderline string
type HTMLInlineCode string
type HTMLCrossed string
type HTMLHyperlink struct {
	Content []HTMLNode
	Link    string
}
type HTMLImage struct {
	Content []HTMLNode
	Path    string
}

func (t HTMLPlain) HTMLRender() string {
	return string(t)
}
func (t HTMLBold) HTMLRender() string {
	return fmt.Sprintf("<b>%s</b>", t)
}
func (t HTMLItalic) HTMLRender() string {
	return fmt.Sprintf("<i>%s</i>", t)
}
func (t HTMLUnderline) HTMLRender() string {
	return fmt.Sprintf("<u>%s</u>", t)
}
func (t HTMLInlineCode) HTMLRender() string {
	return fmt.Sprintf("<code>%s</code>", t)
}
func (t HTMLCrossed) HTMLRender() string {
	return fmt.Sprintf("<strike>%s</strike>", t)
}
func (t HTMLHyperlink) HTMLRender() string {
	return fmt.Sprintf("<a href='%s'>%s</a>", t.Link, htmlRender(t.Content))
}
func (t HTMLImage) HTMLRender() string {
	return fmt.Sprintf("<img src='%s'>%s</img>", t.Path, htmlRender(t.Content))
}

// Containers

type HTMLHeader struct {
	Content []HTMLNode
	Level   int
}
type HTMLParagraph []HTMLNode
type HTMLCode string
type HTMLQuote []HTMLNode
type HTMLBreak bool

type HTMLListItem []HTMLNode
type HTMLUnorderedItem HTMLListItem
type HTMLOrderedItem HTMLListItem
type HTMLOrderedList []HTMLOrderedItem
type HTMLUnorderedList []HTMLUnorderedItem

type HTMLTableItem []HTMLNode
type HTMLTableHeader []TableItem
type HTMLTableRow []TableItem
type HTMLTable struct {
	Header TableHeader
	Rows   []TableRow
}

func (b HTMLHeader) HTMLRender() string {
	return fmt.Sprintf("<h%d>%s</h%d>", b.Level, htmlRender(b.Content), b.Level)
}
func (b HTMLParagraph) HTMLRender() string {
	return fmt.Sprintf("<p>%s</p>", b)
}
func (b HTMLCode) HTMLRender() string {
	return fmt.Sprintf("<pre>%s</pre>", b)
}
func (b HTMLQuote) HTMLRender() string {
	return fmt.Sprintf("<blockquote>%s</blockquote>", b)
}
func (b HTMLBreak) HTMLRender() string {
	return "<br>"
}
func (b HTMLOrderedList) HTMLRender() string {
	result := "<ol>"
	for _, item := range b {
		renderedItem := htmlRender(item)
		result += fmt.Sprintf("<li>%s</li>", renderedItem)
	}
	result += "</ol>"

	return result
}
func (b HTMLUnorderedList) HTMLRender() string {
	result := "<ul>"
	for _, item := range b {
		renderedItem := htmlRender(item)
		result += fmt.Sprintf("<li>%s</li>", renderedItem)
	}
	result += "</ul>"

	return result
}
