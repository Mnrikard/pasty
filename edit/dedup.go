package edit

import (
	"slices"
	"strings"
)

func (e *EditorArgs) Deduplicate(input string) (string, error) {
	items := strings.Split(input, e.RowDelimiter)
	newItems := make([]string, 0)
	for _, item := range items {
		if !slices.Contains(newItems, item) {
			newItems = append(newItems, item)
		}
	}

	replacedText := strings.Join(newItems, e.RowDelimiter)

	return replacedText, nil
}
