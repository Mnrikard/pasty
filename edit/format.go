package edit

import (
	"fmt"
	"strings"

	"github.com/mattr/pasty/reader"
)


func (e *EditorArgs) FormatCode(input string) (string, error) {
	if strings.EqualFold(e.Option, "sql") {
		return formatSql(input)
	}

	return input, fmt.Errorf("No format specified")
}

func formatSql(input string) (string, error) {
	sqlReader := &reader.TextReader {
		InlineComments: "--",
		StartBlockComment: "/*",
		EndBlockComment: "*/",
		StringChar: "'",
		StringEscapeChar: "'",
	}

	sqlReader.SplitCode(input)
	breakBefores := []string { "INNER", "OUTER", "CROSS", "WHERE", "ORDER", "HAVING", "LIMIT", "OFFSET" }
	tabBefore := []string { "AND", "OR" } 
	capitalize := []string { "SELECT", "FROM", "AS", "BY", "AND", "OR", "INNER", "OUTER", "CROSS", "WHERE", "ORDER", "HAVING", "LIMIT", "OFFSET", "JOIN" }
	output := strings.Builder{}

	for i, word := range sqlReader.Words {
		str := word.String()
		if strings.TrimSpace(str) == "" {
			continue
		}

		if has(capitalize, str) {
			str = strings.ToUpper(str)
		}

		var err error

		if has(breakBefores, str) {
			_, err = output.WriteString("\n")
		} else if has(tabBefore, str) {
			_, err = output.WriteString("\n\t")
		} else if i > 0 {
			_, err = output.WriteRune(' ')
		}

		if err != nil {
			return input, err
		}
		
		_, err = output.WriteString(str)
		if err != nil {
			return input, err
		}
	}

	return output.String(), nil
}

func has(list []string, word string) bool {
	for _, lw := range list {
		if strings.EqualFold(lw, word) {
			return true
		}
	}

	return false
}
