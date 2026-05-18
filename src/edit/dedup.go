package edit

import (
	"strings"
)

func (e *EditorArgs) Deduplicate(input string) (string, error) {
	items := strings.Split(input, e.RowDelimiter)
	set := make(map[string]bool, 0)
	output := make([]string, 0)
	for _, item := range items {
		_, exists := set[item]
		if !exists {
			output = append(output, item)
		}
		set[item] = true
	}

	return strings.Join(output, e.RowDelimiter), nil
}
