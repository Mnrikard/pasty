package edit

import "testing"

func TestUpper(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"lowercase to upper", "hello world", "HELLO WORLD"},
		{"mixed case", "HeLLo WoRLD", "HELLO WORLD"},
		{"already upper", "HELLO", "HELLO"},
		{"with numbers", "hello123", "HELLO123"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{}
			result, err := args.Upper(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}

func TestLower(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"uppercase to lower", "HELLO WORLD", "hello world"},
		{"mixed case", "HeLLo WoRLD", "hello world"},
		{"already lower", "hello", "hello"},
		{"with numbers", "HELLO123", "hello123"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{}
			result, err := args.Lower(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}

func TestTitle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"lowercase words", "hello world", "Hello World"},
		{"all caps", "HELLO WORLD", "Hello World"},
		{"mixed case", "hELLO wORLD", "Hello World"},
		{"single word", "hello", "Hello"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{}
			result, err := args.Title(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}
