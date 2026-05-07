package util

import (
	"reflect"
	"testing"
)

func TestSplitRows(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"line1\nline2\nline3", []string{"line1", "line2", "line3"}},
		{"line1\r\nline2\r\nline3", []string{"line1", "line2", "line3"}},
		{"", []string{""}},
	}

	for _, tt := range tests {
		result := SplitRows(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("SplitRows(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

func TestSplitColumns(t *testing.T) {
	rows := []string{"a,b,c", "d,e,f"}
	expected := [][]string{{"a", "b", "c"}, {"d", "e", "f"}}
	result := SplitColumns(rows, ",")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SplitColumns() = %v; want %v", result, expected)
	}
}

func TestIsNullOrNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123", true},
		{"12.3", true},
		{"NULL", true},
		{"null", true},
		{"abc", false},
		{"", false},
	}

	for _, tt := range tests {
		result := IsNullOrNumber(tt.input)
		if result != tt.expected {
			t.Errorf("IsNullOrNumber(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

func TestEscapeRegex(t *testing.T) {
	input := "().+*-|"
	expected := "\\(\\)\\.\\+\\*\\-\\|"
	result := EscapeRegex(input)
	if result != expected {
		t.Errorf("EscapeRegex(%q) = %q; want %q", input, result, expected)
	}
}
