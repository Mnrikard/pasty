package edit

import (
	"regexp"
)

func (e *EditorArgs) ReplaceText(input string) (string, error) {
	e.PrependRegex()
	rx, err := regexp.Compile(e.Regex)
	if err != nil {
		return input, err
	}

	replacedText := rx.ReplaceAllString(input, e.Replacement)

	return replacedText, nil
}
