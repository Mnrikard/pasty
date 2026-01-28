package edit

import "testing"

func TestCountItem_Words(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		option string
	}{
		{"single word", "hello", "words"},
		{"multiple words", "hello world foo", "words"},
		{"word option", "one two", "word"},
		{"extra spaces", "  hello   world  ", "words"},
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
