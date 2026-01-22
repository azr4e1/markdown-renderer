package main

type Text string
type Bold Text
type Italic Text
type Underline Text
type Crossed Text

type Header1 string
type Header2 string
type Header3 string
type Header4 string
type Header5 string
type Header6 string

type ListItem string
type UnorderedItem ListItem
type OrderedItem ListItem
type List []ListItem

type Hyperlink struct {
	Content string
	Link    string
}

type Image struct {
	Content string
	Path    string
}

type TableHeader []Text
type TableRow []Text
type Table struct {
	Header TableHeader
	Rows   []TableRow
}
