package edit

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

type sortEntry struct {
	OriginalValue string
	DateValue *time.Time
	FloatValue *float64
}

func (e *EditorArgs) Sort(input string) (string, error) {
	items := strings.Split(input, e.RowDelimiter)
	sortable := getSortEntries(items)
	sort.Slice(sortable, func(a, b int) bool {
		if sortable[a].DateValue != nil && sortable[b].DateValue != nil {
			return sortable[a].DateValue.UnixMicro() > sortable[b].DateValue.UnixMicro()
		}

		if sortable[a].FloatValue != nil && sortable[b].FloatValue != nil {
			return *sortable[a].FloatValue > *sortable[b].FloatValue
		}

		if e.Option == "i" {
			return strings.ToLower(sortable[a].OriginalValue) > strings.ToLower(sortable[b].OriginalValue)
		}

		if e.Switches.Invert {
			return sortable[a].OriginalValue > sortable[b].OriginalValue
		}

		return sortable[a].OriginalValue < sortable[b].OriginalValue
	})

	output := make([]string, len(items))
	for i := range sortable {
		output[i] = sortable[i].OriginalValue
	}

	return strings.Join(output, e.RowDelimiter), nil
}

func getSortEntries(items []string) []sortEntry {
	output := make([]sortEntry, len(items))
	for i, item := range items {
		output[i] = sortEntry {
			OriginalValue: item,
		}

		date := getDateValue(item)
		if date != nil {
			output[i].DateValue = date
		}

		flt, err := strconv.ParseFloat(item, 64)
		if err == nil {
			output[i].FloatValue = &flt
		}
	}

	return output
}

var standardDateFormats = []string {
	time.RFC3339,
	"2006-01-02 15:04:05",
	"2006-01-02 03:04:05",
	"2006-01-02 3:04:05",
	"2006-01-02 3:04:05 PM",
	"01-02-2006 15:04:05",
	"01-02-2006 03:04:05",
	"01-02-2006 3:04:05",
	"01-02-2006 3:04:05 PM",
	"01-02-2006",
	"2006-01-02",
	"15:04:05",
	"03:04:05",
	"03:04:05 PM",
}

func getDateValue(input string) *time.Time {
	for _, fmt := range standardDateFormats {
		date, err := time.Parse(fmt, input)
		if err == nil {
			return &date
		}
	}

	return nil
}

