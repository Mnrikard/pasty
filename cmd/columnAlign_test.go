package cmd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mattr/pasty/text"
)

func TestColumnsByTab(t *testing.T) {
	text.SetMockedText(`one	two	three
four	five	six
seven	eight	nine
ten	eleven	twelve
thirteen	fourteen	fifteen`, nil)
	expectedOutput := "" +
		"one       two       three    \n" +
		"four      five      six      \n" +
		"seven     eight     nine     \n" +
		"ten       eleven    twelve   \n" +
		"thirteen  fourteen  fifteen  "
	ColumnAlign.Command.Run(nil, []string { })
	assertEqual(t, expectedOutput, text.GetMockedText(), "")
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
