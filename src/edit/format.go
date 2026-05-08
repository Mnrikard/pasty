package edit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Mnrikard/pasty/reader"
)

func (e *EditorArgs) FormatCode(input string) (string, error) {
	if strings.EqualFold(e.Option, "sql") {
		return formatSql(input)
	}
	if strings.EqualFold(e.Option, "json") {
		//try to format it as valid json first
		jsonBytes := &bytes.Buffer{}
		if json.Indent(jsonBytes, []byte(input), "", settings().TabString) != nil {
			//then fall back to my imperfect json parser
			return formatJson(input)
		}

		return jsonBytes.String(), nil
	}

	return input, fmt.Errorf("No format specified")
}

func formatSql(input string) (string, error) {
	sqlReader := &reader.TextReader{
		InlineComments:    "--",
		StartBlockComment: "/*",
		EndBlockComment:   "*/",
		StringChars:       []string{"'"},
		StringEscapeChar:  "'",
	}

	sqlReader.SplitCode(input)
	breakBefores := []string{"INNER", "OUTER", "CROSS", "WHERE", "ORDER", "HAVING", "LIMIT", "OFFSET"}
	tabBefore := []string{"AND", "OR"}
	capitalize := []string{"SELECT", "FROM", "AS", "BY", "AND", "OR", "INNER", "OUTER", "CROSS", "WHERE", "ORDER", "HAVING", "LIMIT", "OFFSET", "JOIN"}
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
			_, err = output.WriteString(fmt.Sprintf("\n%v", settings().TabString))
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

func formatJson(input string) (string, error) {
	jsonReader := &reader.TextReader{
		InlineComments:    "//",
		StartBlockComment: "/*",
		EndBlockComment:   "*/",
		StringChars:       []string{"'", "\"", "`"},
		StringEscapeChar:  "'",
		KeyWords:          []string{":", "}", "{", ",", "[", "]"},
	}

	jsonReader.SplitCode(input)
	dedentBefores := []string{"}", "]"}
	indentAfters := []string{"{", "["}
	breakAfters := []string{","}
	output := strings.Builder{}
	tabCount := 0
	tabStr := ""

	for _, word := range jsonReader.Words {
		str := word.String()
		if strings.TrimSpace(str) == "" {
			continue
		}

		var err error

		if has(dedentBefores, str) {
			tabCount--
			tabStr = tab(tabCount)
			_, err = output.WriteString(fmt.Sprintf("\n%v%v", tabStr, str))
		} else if has(indentAfters, str) {
			tabCount++
			tabStr = tab(tabCount)
			_, err = output.WriteString(fmt.Sprintf("%v\n%v", str, tabStr))
		} else if has(breakAfters, str) {
			_, err = output.WriteString(fmt.Sprintf("%v\n%v", str, tabStr))
		} else if str == ":" {
			_, err = output.WriteString(": ")
		} else {
			_, err = output.WriteString(str)
		}

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

func tab(size int) string {
	output := make([]string, size)
	for i := range output {
		output[i] = settings().TabString
	}

	return strings.Join(output, "")
}
