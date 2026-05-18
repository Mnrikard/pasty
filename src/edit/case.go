package edit

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (e *EditorArgs) Upper(input string) (string, error) {
	return strings.ToUpper(input), nil
}

func (e *EditorArgs) Lower(input string) (string, error) {
	return strings.ToLower(input), nil
}

func (e *EditorArgs) Title(input string) (string, error) {
	caser := cases.Title(language.English)
	return caser.String(input), nil
}
