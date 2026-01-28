package main

type Node interface {
	ToHTML() HTMLNode
}

func markdownToHTML(nodes []Node) []HTMLNode {
	htmlNodes := []HTMLNode{}

	for _, n := range nodes {
		htmlNodes = append(htmlNodes, n.ToHTML())
	}

	return htmlNodes
}

// Leaves

type Plain string
type Bold string
type Italic string
type Underline string
type InlineCode string
type Crossed string
type Hyperlink struct {
	Content []Node
	Link    string
}
type Image struct {
	Content []Node
	Path    string
}

func (t Plain) ToHTML() HTMLNode {
	return HTMLPlain(t)
}
func (t Bold) ToHTML() HTMLNode {
	return HTMLBold(t)
}
func (t Italic) ToHTML() HTMLNode {
	return HTMLItalic(t)
}
func (t Underline) ToHTML() HTMLNode {
	return HTMLUnderline(t)
}
func (t InlineCode) ToHTML() HTMLNode {
	return HTMLInlineCode(t)
}
func (t Crossed) ToHTML() HTMLNode {
	return HTMLCrossed(t)
}
func (t Hyperlink) ToHTML() HTMLNode {
	return HTMLHyperlink{Content: markdownToHTML(t.Content), Link: t.Link}
}
func (t Image) ToHTML() HTMLNode {
	return HTMLImage{Content: markdownToHTML(t.Content), Path: t.Path}
}

// Containers

type Header struct {
	Content []Node
	Level   int
}
type Paragraph []Node
type Code string
type Quote []Node
type Break bool

type ListItem []Node
type UnorderedItem ListItem
type OrderedItem ListItem
type OrderedList []OrderedItem
type UnorderedList []UnorderedItem

type TableItem []Node
type TableHeader []TableItem
type TableRow []TableItem
type Table struct {
	Header TableHeader
	Rows   []TableRow
}

func (b Header) ToHTML() HTMLNode {
	return nil
}
func (b Paragraph) ToHTML() HTMLNode {
	return nil
}
func (b Code) ToHTML() HTMLNode {
	return nil
}
func (b Quote) ToHTML() HTMLNode {
	return nil
}
func (b Break) ToHTML() HTMLNode {
	return nil
}
func (b OrderedList) ToHTML() HTMLNode {
	return nil
}
func (b UnorderedList) ToHTML() HTMLNode {
	return nil
}
func (b Table) ToHTML() HTMLNode {
	return nil
}
