package edit

import (
	"fmt"
	"testing"
)

func TestCountItem_Words(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		option        string
		expectedCount string
	}{
		{"single word", "hello", "words", "1 words"},
		{"multiple words", "hello world foo", "words", "3 words"},
		{"word option", "one two", "word", "2 word"},
		{"extra spaces", "  hello   world  ", "words", "2 words"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{Option: tt.option}
			actualCount := ""
			notify = func(input string) {
				actualCount = input
			}
			_, err := args.CountItem(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expectedCount, actualCount, fmt.Sprintf("Expected %q, but was %q", tt.expectedCount, actualCount))
		})
	}
}

func TestCountItem_Lines(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		option string
	}{
		{"single line", "hello", "lines"},
		{"multiple lines", "hello\nworld\nfoo", "lines"},
		{"line option", "one\ntwo", "line"},
		{"windows newlines", "hello\r\nworld", "lines"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{Option: tt.option}
			result, err := args.CountItem(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.input, result, "CountItem should return original input")
		})
	}
}

func TestCountItem_Characters(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"simple text", "hello"},
		{"with spaces", "hello world"},
		{"empty after trim", "   "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{Option: "chars"}
			result, err := args.CountItem(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.input, result, "CountItem should return original input")
		})
	}
}
