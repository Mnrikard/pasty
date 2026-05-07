package edit

import (
	"testing"

	"github.com/Mnrikard/pasty/util"
)

func testSettings() *util.SettingValues {
	return &util.SettingValues{
		TabString: "\t",
		DateFormats: []string{
			"2006-01-02",
			"2006-1-2",
			"01-02-2006",
			"1-2-2006",
			"2006/01/02",
			"2006/1/2",
			"01/02/2006",
			"1/2/2006",
		},
		TimeFormats: []string{
			"15:04:05",
			"03:04:05",
			"3:04:05",
			"3:04:05 PM",
		},
	}
}


func TestFormatCode_Sql(t *testing.T) {
	settings = testSettings
	input := "select * from table where id = 1 and name = 'test'"
	// formatSql logic:
	// SELECT (capitalize)
	// *
	// FROM (capitalize)
	// table
	// WHERE (capitalize, break before) -> \nWHERE
	// id
	// =
	// 1
	// AND (capitalize, tab before) -> \n\tAND
	// name
	// =
	// 'test'
	expected := "SELECT * FROM table\nWHERE id = 1\n\tAND name = 'test'"

	args := EditorArgs{Option: "sql"}
	result, err := args.FormatCode(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, expected, result, "")
}

func TestFormatCode_Json(t *testing.T) {
	settings = testSettings
	input := `{"a":1,"b":[1,2]}`
	// Trace:
	// { -> indentAfter -> {\n\t
	// "a" -> "a"
	// : -> : 
	// 1 -> 1
	// , -> breakAfter -> ,\n\t
	// "b" -> "b"
	// : -> : 
	// [ -> indentAfter -> [\n\t\t
	// 1 -> 1
	// , -> breakAfter -> ,\n\t\t
	// 2 -> 2
	// ] -> dedentBefore -> \n\t]
	// } -> dedentBefore -> \n}
	expected := "{\n\t\"a\": 1,\n\t\"b\": [\n\t\t1,\n\t\t2\n\t]\n}"

	args := EditorArgs{Option: "json"}
	result, err := args.FormatCode(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, expected, result, "")
}

func TestFormatCode_NoFormat(t *testing.T) {
	settings = testSettings
	args := EditorArgs{Option: "invalid"}
	input := "some code"
	result, err := args.FormatCode(input)
	if err == nil {
		t.Fatal("expected error for invalid format")
	}
	if result != input {
		t.Errorf("expected input to be returned on error, got %q", result)
	}
}
