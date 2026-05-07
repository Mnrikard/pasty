package edit

import "testing"

func TestGrep_MatchingLines(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		regex        string
		rowDelimiter string
		expected     string
	}{
		{
			"simple match",
			"apple\nbanana\napricot",
			"^a",
			"\n",
			"apple\napricot",
		},
		{
			"no matches",
			"banana\ncherry",
			"^a",
			"\n",
			"",
		},
		{
			"all match",
			"apple\napricot",
			"^a",
			"\n",
			"apple\napricot",
		},
		{
			"case sensitive",
			"Apple\napple",
			"^a",
			"\n",
			"apple",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{
				Regex:        tt.regex,
				RowDelimiter: tt.rowDelimiter,
				Invert:       false,
			}
			result, err := args.Grep(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}

func TestGrep_InvertMatch(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		regex        string
		rowDelimiter string
		expected     string
	}{
		{
			"invert simple",
			"apple\nbanana\napricot",
			"^a",
			"\n",
			"banana",
		},
		{
			"invert all match",
			"apple\napricot",
			"^a",
			"\n",
			"",
		},
		{
			"invert no match",
			"banana\ncherry",
			"^a",
			"\n",
			"banana\ncherry",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{
				Regex:        tt.regex,
				RowDelimiter: tt.rowDelimiter,
				Invert:       true,
			}
			result, err := args.Grep(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}

func TestGrep_OnlyMatching(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		regex        string
		rowDelimiter string
		expected     string
	}{
		{
			"extract numbers",
			"abc123def456",
			"[0-9]+",
			"\n",
			"123\n456",
		},
		{
			"extract words",
			"hello world foo",
			"[a-z]+",
			",",
			"hello,world,foo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{
				Regex:        tt.regex,
				RowDelimiter: tt.rowDelimiter,
				Option:       "OnlyMatching",
			}
			result, err := args.Grep(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}

func TestGrep_InvalidRegex(t *testing.T) {
	args := EditorArgs{
		Regex:        "[invalid",
		RowDelimiter: "\n",
	}
	_, err := args.Grep("test input")
	if err == nil {
		t.Fatal("expected error for invalid regex")
	}
}
