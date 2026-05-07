package edit

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Mnrikard/pasty/util"
)

func (e *EditorArgs) CountItem(input string) (string, error) {
	var count int
	var rx *regexp.Regexp

	switch e.Option {
	case "words","word":
		rx = regexp.MustCompile(`\s+`)
		count = len(rx.Split(strings.TrimSpace(input), -1))
	case "lines","line":
		rx = regexp.MustCompile("\r?\n")
		count = len(rx.Split(strings.TrimSpace(input), -1))
	default:
		count = len(strings.TrimSpace(input))
	}

	util.Notify(fmt.Sprintf("%d %s", count, e.Option))

	return input, nil
}
