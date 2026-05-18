package edit

import (
	"strings"
	"testing"
)

func TestInsertSQL(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		columnDelimiter string
		rowDelimiter    string
		option          string
		expectContains  []string
	}{
		{
			"basic insert",
			"id\tname\n1\tAlice\n2\tBob",
			"\t",
			"\n",
			"users",
			[]string{"insert into [users]", "(id, name)", "(1, 'Alice')", "(2, 'Bob')"},
		},
		{
			"with schema",
			"id\tvalue\n1\ttest",
			"\t",
			"\n",
			"dbo.mytable",
			[]string{"insert into [dbo].[mytable]"},
		},
		{
			"numeric values",
			"id\tamount\n1\t100",
			"\t",
			"\n",
			"transactions",
			[]string{"(1, 100)"},
		},
		{
			"string with apostrophe",
			"id\tname\n1\tO'Brien",
			"\t",
			"\n",
			"users",
			[]string{"'O''Brien'"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := EditorArgs{
				ColumnDelimiter: tt.columnDelimiter,
				RowDelimiter:    tt.rowDelimiter,
				Option:          tt.option,
			}
			result, err := args.InsertSQL(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			for _, expected := range tt.expectContains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected result to contain %q, got %q", expected, result)
				}
			}
		})
	}
}

func TestInsertSQL_SkipsEmptyLines(t *testing.T) {
	input := "id\tname\n1\tAlice\n\n2\tBob\n"
	args := EditorArgs{
		ColumnDelimiter: "\t",
		RowDelimiter:    "\n",
		Option:          "users",
	}
	result, err := args.InsertSQL(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if strings.Contains(result, "()") {
		t.Error("expected empty lines to be skipped")
	}
}
