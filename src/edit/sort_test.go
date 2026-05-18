package edit

import (
	"testing"

	"github.com/Mnrikard/pasty/switches"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		rowDelimiter string
		option       string
		expected     string
	}{
		{
			"sorts strings descending",
			"banana\napple\ncherry",
			"\n",
			"",
			"cherry\nbanana\napple",
		},
		{
			"sorts strings case insensitive",
			"banana\nApple\ncherry",
			"\n",
			"i",
			"cherry\nbanana\nApple",
		},
		{
			"sorts numbers descending",
			"10\n2\n100\n3",
			"\n",
			"",
			"100\n10\n3\n2",
		},
		{
			"sorts float numbers descending",
			"3.14\n1.5\n2.71\n0.99",
			"\n",
			"",
			"3.14\n2.71\n1.5\n0.99",
		},
		{
			"sorts dates descending",
			"2024-01-15\n2023-06-01\n2025-12-31",
			"\n",
			"",
			"2025-12-31\n2024-01-15\n2023-06-01",
		},
		{
			"sorts datetime descending",
			"2024-01-15 10:30:00\n2024-01-15 08:00:00\n2024-01-15 23:59:59",
			"\n",
			"",
			"2024-01-15 23:59:59\n2024-01-15 10:30:00\n2024-01-15 08:00:00",
		},
		{
			"single item",
			"apple",
			"\n",
			"",
			"apple",
		},
		{
			"comma delimiter",
			"banana,apple,cherry",
			",",
			"",
			"cherry,banana,apple",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sw := switches.Switches{Invert: true}
			args := EditorArgs{RowDelimiter: tt.rowDelimiter, Option: tt.option, Switches: &sw}
			result, err := args.Sort(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertEqual(t, tt.expected, result, "")
		})
	}
}
