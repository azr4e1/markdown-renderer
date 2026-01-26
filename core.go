package main

import "fmt"

// Inline text
type Text interface {
	isText() bool
}

type Plain string
type Bold string
type Italic string
type Underline string
type InlineCode string
type Crossed string
type Hyperlink struct {
	Content Text
	Link    string
}
type Image struct {
	Content Text
	Path    string
}

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
func (t Hyperlink) isText() bool {
	return true
}
func (t Image) isText() bool {
	return true
}

// Markdown block
type Block interface {
	isBlock() bool
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

func (t Plain) String() string {
	return fmt.Sprintf("Plain(%s)", string(t))
}
func (t Bold) String() string {
	return fmt.Sprintf("Bold(%s)", string(t))
}
func (t Italic) String() string {
	return fmt.Sprintf("Italic(%s)", string(t))
}
func (t Underline) String() string {
	return fmt.Sprintf("Underline(%s)", string(t))
}
func (t InlineCode) String() string {
	return fmt.Sprintf("InlineCode(%s)", string(t))
}
func (t Crossed) String() string {
	return fmt.Sprintf("Crossed(%s)", string(t))
}
