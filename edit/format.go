package edit

import (
	"fmt"
	"strings"
)


func (e *EditorArgs) FormatCode(input string) (string, error) {
	if strings.EqualFold(e.Option, "sql") {
		return formatSql(input)
	}

	return input, fmt.Errorf("No format specified")
}

func formatSql(input string) (string, error) {
	panic("not implemented")
}
