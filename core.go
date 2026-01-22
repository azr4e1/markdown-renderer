package main

type Text string
type Bold Text
type Italic Text
type Underline Text
type Crossed Text

type Header struct {
	Content []Text
	Level   int
}

type Code Text
type Quote []Text

type ListItem []Text
type UnorderedItem ListItem
type OrderedItem ListItem
type List []ListItem

type Hyperlink struct {
	Content Text
	Link    string
}

type Image struct {
	Content Text
	Path    string
}

type TableItem []Text
type TableHeader []TableItem
type TableRow []TableItem
type Table struct {
	Header TableHeader
	Rows   []TableRow
}
