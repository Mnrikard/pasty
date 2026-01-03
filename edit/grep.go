package edit

import (
	"regexp"
	"strings"
)

func (e *EditorArgs) Grep(input string) (string, error) {
	rx, err := regexp.Compile(e.Regex)
	if err != nil {
		return input, err
	}

	if e.Option == "L" {
		matches := rx.FindAllString(input, -1)
		return strings.Join(matches, e.RowDelimiter), nil
	}

	output := make([]string, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if rx.MatchString(line) {
			output = append(output, strings.Trim(line, "\r"))
		}
	}

	return strings.Join(output, e.RowDelimiter), nil

}
