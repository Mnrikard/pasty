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
		//Sort Dates
		if sortable[a].DateValue != nil && sortable[b].DateValue != nil {
			if e.Switches.Invert {
				return sortable[a].DateValue.UnixMicro() > sortable[b].DateValue.UnixMicro()
			}
			return sortable[a].DateValue.UnixMicro() < sortable[b].DateValue.UnixMicro()
		}

		//Sort Numbers
		if sortable[a].FloatValue != nil && sortable[b].FloatValue != nil {
			if e.Switches.Invert {
				return *sortable[a].FloatValue > *sortable[b].FloatValue
			}
			return *sortable[a].FloatValue < *sortable[b].FloatValue
		}

		//sort as strings
		if sortable[a].DateValue == nil && sortable[b].DateValue == nil && sortable[a].FloatValue == nil && sortable[b].FloatValue == nil {
			if e.Option == "i" {
				if e.Switches.Invert {
					return strings.ToLower(sortable[a].OriginalValue) > strings.ToLower(sortable[b].OriginalValue)
				}
				return strings.ToLower(sortable[a].OriginalValue) < strings.ToLower(sortable[b].OriginalValue)
			}

			if e.Switches.Invert {
				return sortable[a].OriginalValue > sortable[b].OriginalValue
			}

			return sortable[a].OriginalValue < sortable[b].OriginalValue
		}

		//numbers go on top
		if sortable[a].FloatValue != nil {
			return !e.Switches.Invert
		}
		if sortable[b].FloatValue != nil {
			return e.Switches.Invert
		}

		//dates go before strings
		if sortable[a].DateValue != nil {
			return !e.Switches.Invert
		}
		if sortable[b].DateValue != nil {
			return e.Switches.Invert
		}

		//should never reach this
		return false
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
	"2006-01-02",
	"01-02-2006",
}

var standardTimeFormats = []string {
	"15:04:05",
	"03:04:05",
	"3:04:05",
	"3:04:05 PM",
}

func getDateValue(input string) *time.Time {
	for _, dfmt := range standardDateFormats {
		date, err := time.Parse(dfmt, input)
		if err == nil {
			return &date
		}
		for _, tfmt := range standardTimeFormats {
			date, err = time.Parse(dfmt + " " + tfmt, input)
			if err == nil {
				return &date
			}
			date, err = time.Parse(dfmt + "T" + tfmt, input)
			if err == nil {
				return &date
			}
		}
	}

	for _, tfmt := range standardTimeFormats {
		date, err := time.Parse(tfmt, input)
		if err == nil {
			return &date
		}
	}

	return nil
}

