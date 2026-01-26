package main

type Text interface {
	isText() bool
}

type Plain string
type Bold string
type Italic string
type Underline string
type InlineCode string
type Crossed string

func (t Plain) isText() bool {
	return true
}
func (t Bold) isText() bool {
	return true
}
func (t Italic) isText() bool {
	return true
}
func (t Underline) isText() bool {
	return true
}
func (t InlineCode) isText() bool {
	return true
}
func (t Crossed) isText() bool {
	return true
}

type Hyperlink struct {
	Content Text
	Link    string
}

type Image struct {
	Content Text
	Path    string
}

type Header struct {
	Content []Text
	Level   int
}
type Paragraph []Text
type Code string
type Quote []Text
type Break bool

type ListItem []Text
type UnorderedItem ListItem
type OrderedItem ListItem
type OrderedList []OrderedItem
type UnorderedList []UnorderedItem

type TableItem []Text
type TableHeader []TableItem
type TableRow []TableItem
type Table struct {
	Header TableHeader
	Rows   []TableRow
}

type Block interface {
	isBlock() bool
}

func (b Header) isBlock() bool {
	return true
}
func (b Paragraph) isBlock() bool {
	return true
}
func (b Code) isBlock() bool {
	return true
}
func (b Quote) isBlock() bool {
	return true
}
func (b Break) isBlock() bool {
	return true
}
func (b OrderedList) isBlock() bool {
	return true
}
func (b UnorderedList) isBlock() bool {
	return true
}
func (b Table) isBlock() bool {
	return true
}
