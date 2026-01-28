package edit

import (
	"regexp"
	"testing"
)

func TestSetText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		option   string
		expected string
	}{
		{"replace with text", "original", "replacement", "replacement"},
		{"empty replacement", "original", "", ""},
		{"empty input", "", "replacement", "replacement"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{Option: tt.option}
			result, err := args.SetText(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}

func TestNewGuid(t *testing.T) {
	args := EditorArgs{}
	result, err := args.NewGuid("ignored input")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	uuidPattern := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	if !uuidPattern.MatchString(result) {
		t.Fatalf("expected UUID format, got %q", result)
	}
}

func TestNewGuid_Unique(t *testing.T) {
	args := EditorArgs{}
	guid1, _ := args.NewGuid("")
	guid2, _ := args.NewGuid("")

	if guid1 == guid2 {
		t.Fatal("expected unique GUIDs, got duplicates")
	}
}
