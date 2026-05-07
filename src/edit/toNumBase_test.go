package edit

import "testing"

func TestToNumBase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		option   string
		expected string
	}{
		{"decimal to hex", "255", "16", "ff"},
		{"decimal to binary", "10", "2", "1010"},
		{"decimal to octal", "64", "8", "100"},
		{"0x option for hex", "16", "0x", "10"},
		{"default to octal", "8", "", "10"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{Option: tt.option}
			result, err := args.ToNumBase(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}

func TestToNumBase_InvalidInput(t *testing.T) {
	args := EditorArgs{Option: "16"}
	_, err := args.ToNumBase("not a number")
	if err == nil {
		t.Fatal("expected error for invalid input")
	}
}

func TestFromNumBase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		option   string
		expected string
	}{
		{"hex to decimal", "ff", "16", "255"},
		{"binary to decimal", "1010", "2", "10"},
		{"octal to decimal", "100", "8", "64"},
		{"0x option for hex", "10", "0x", "16"},
		{"default from octal", "10", "", "8"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{Option: tt.option}
			result, err := args.FromNumBase(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}

func TestFromNumBase_InvalidInput(t *testing.T) {
	args := EditorArgs{Option: "16"}
	_, err := args.FromNumBase("xyz")
	if err == nil {
		t.Fatal("expected error for invalid input")
	}
}
