package edit

import "testing"

func TestReplaceText(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		regex       string
		replacement string
		expected    string
	}{
		{
			"simple replace",
			"hello world",
			"world",
			"universe",
			"hello universe",
		},
		{
			"replace all occurrences",
			"cat cat cat",
			"cat",
			"dog",
			"dog dog dog",
		},
		{
			"regex pattern",
			"abc123def456",
			"[0-9]+",
			"#",
			"abc#def#",
		},
		{
			"capture groups",
			"hello world",
			"(hello) (world)",
			"$2 $1",
			"world hello",
		},
		{
			"no match",
			"hello world",
			"xyz",
			"abc",
			"hello world",
		},
		{
			"empty replacement",
			"hello world",
			" world",
			"",
			"hello",
		},
		{
			"replace with special chars",
			"hello world",
			"world",
			"wor$$ld",
			"hello wor$ld",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{
				Regex:       tt.regex,
				Replacement: tt.replacement,
			}
			result, err := args.ReplaceText(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}

func TestReplaceText_InvalidRegex(t *testing.T) {
	args := EditorArgs{
		Regex:       "[invalid",
		Replacement: "test",
	}
	result, err := args.ReplaceText("test input")
	if err == nil {
		t.Fatal("expected error for invalid regex")
	}
	assertEqual(t, "test input", result, "should return original on error")
}
