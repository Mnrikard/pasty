package edit

import "testing"

func TestDeduplicate(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		rowDelimiter string
		expected     string
	}{
		{
			"remove duplicate lines",
			"apple\nbanana\napple\ncherry",
			"\n",
			"apple\nbanana\ncherry",
		},
		{
			"no duplicates",
			"apple\nbanana\ncherry",
			"\n",
			"apple\nbanana\ncherry",
		},
		{
			"all duplicates",
			"apple\napple\napple",
			"\n",
			"apple",
		},
		{
			"comma delimiter",
			"a,b,a,c,b",
			",",
			"a,b,c",
		},
		{
			"single item",
			"apple",
			"\n",
			"apple",
		},
		{
			"empty string",
			"",
			"\n",
			"",
		},
		{
			"preserve order",
			"cherry\napple\nbanana\napple",
			"\n",
			"cherry\napple\nbanana",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{RowDelimiter: tt.rowDelimiter}
			result, err := args.Deduplicate(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}
