package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLineParser(t *testing.T) {
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
			name:     "underline with tilde",
			input:    "~underline~",
			expected: []Text{Underline("underline")},
		},
		{
			name:     "inline code with backtick",
			input:    "`code`",
			expected: []Text{InlineCode("code")},
		},
		{
			name:     "crossed with dash",
			input:    "-crossed-",
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
			input:    "*Ci ao c om*-ok -**e va**",
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
			input:    "-abc- ~def~",
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
			got := LineParser(tt.input)
			if diff := cmp.Diff(got, tt.expected); diff != "" {
				t.Errorf("LineParser(%q)\n  got:      %v\n  expected: %v\n  Diff:      %s", tt.input, got, tt.expected, diff)
			}
		})
	}
}
