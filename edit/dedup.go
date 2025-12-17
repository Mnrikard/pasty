package edit

import (
	"log"
	"slices"
	"strings"
)

func (e *EditorArgs) Deduplicate(input string) (string, error) {
	items := strings.Split(input, e.RowDelimiter)
	newItems := make([]string, 0)
	for _, item := range items {
		log.Printf("found %s\n", item)
		if !slices.Contains(newItems, item) {
			log.Println("added")
			newItems = append(newItems, item)
		}
	}

	replacedText := strings.Join(newItems, e.RowDelimiter)

	return replacedText, nil
}
