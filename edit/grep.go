package edit

import (
	"regexp"
	"strings"
)

func (e *EditorArgs) Grep(input string) (string, error) {
	e.PrependRegex()
	rx, err := regexp.Compile(e.Regex)
	if err != nil {
		return input, err
	}

	if e.Switches != nil && e.Switches.GrepOnlyMatching {
		matches := rx.FindAllString(input, -1)
		return strings.Join(matches, e.RowDelimiter), nil
	}

	invert := e.Switches != nil && e.Switches.Invert

	output := make([]string, 0)
	lines := strings.SplitSeq(input, e.RowDelimiter)
	for line := range lines {
		if rx.MatchString(line) {
			if !invert {
				output = append(output, strings.Trim(line, "\r"))
			}
		} else {
			if invert {
				output = append(output, strings.Trim(line, "\r"))
			}
		}
	}

	return strings.Join(output, e.RowDelimiter), nil
}
