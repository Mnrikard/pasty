package edit

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Mnrikard/pasty/util"
)

func (e *EditorArgs) InsertSQL(input string) (string, error) {
	rowOfInsert := 1000
	lineSplitter := regexp.MustCompile(e.RowDelimiter)
	lines := lineSplitter.Split(input, -1)
	columnNames := e.getColumns(lines[0])
	insertStatement := e.defineInsertStatement(columnNames)

	output := strings.Builder{}

	for _, line := range lines[1:] {
		if strings.Trim(line, "\t\r\n ") == "" {
			continue
		}

		if rowOfInsert >= 1000 {
			output.WriteString(insertStatement)
			rowOfInsert = 0
		} else {
			output.WriteString(",")
		}

		cols := e.getColumns(line)
		writeSingleRow(cols, &output)
		rowOfInsert = rowOfInsert + 1
	}

	return output.String(), nil
}

func (e *EditorArgs) defineInsertStatement(columns []string) string {
	tableName := strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(e.Option,
					".", "].["),
				"[[", "["),
			"]]", "]"),
		"[]", "")
	return fmt.Sprintf("insert into [%s] (%s)\nvalues\n ", tableName, strings.Join(columns, ", "))
}

func (e *EditorArgs) getColumns(line string) []string {
	colSplitter := regexp.MustCompile(e.ColumnDelimiter)
	return colSplitter.Split(line, -1)
}

func writeSingleRow(cols []string, b *strings.Builder) {
	b.WriteString("(")
	for j, col := range cols {
		if j > 0 {
			b.WriteString(", ")
		}

		var apostropheOrNot string
		if util.IsNullOrNumber(col) {
			apostropheOrNot = ""
		} else {
			apostropheOrNot = "'"
		}

		fmt.Fprintf(b, "%s%s%s", apostropheOrNot, strings.ReplaceAll(col, "'", "''"), apostropheOrNot)
	}

	b.WriteString(")\n")
}
