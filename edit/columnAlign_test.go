package edit

import (
	"fmt"
	"strings"
	"testing"
)

func TestColumnsByTab(t *testing.T) {
	input := `one	two	three
four	five	six
seven	eight	nine
ten	eleven	twelve
thirteen	fourteen	fifteen`
	expectedOutput := "" +
		"one       two       three  \n" +
		"four      five      six    \n" +
		"seven     eight     nine   \n" +
		"ten       eleven    twelve \n" +
		"thirteen  fourteen  fifteen"
	args := EditorArgs{
		NumSpaces: 2,
		ColumnDelimiter: "\t",
	}
	actualOutput, _ := args.AlignColumns(input)
	assertEqual(t, expectedOutput, actualOutput, "")
}

func TestEmptyInputProducesEmptyOutput(t *testing.T) {
	input := ""
	expectedOutput := ""
	args := EditorArgs {
		NumSpaces: 2,
		ColumnDelimiter: "\t",
	}

	actualOutput, _ := args.AlignColumns(input)
	assertEqual(t, expectedOutput, actualOutput, "")
}

func assertEqual(t *testing.T, expected, actual string, message string) {
	if expected == actual {
		return
	}
	if len(message) == 0 {
		expected = strings.ReplaceAll(strings.ReplaceAll(expected, "\r","\\r"), "\t", "\\t")
		actual = strings.ReplaceAll(strings.ReplaceAll(actual, "\r","\\r"), "\t", "\\t")

		message = fmt.Sprintf("expected %q but was %q", expected, actual)
	}

	t.Fatal(message)
}
