package edit

import (
	"bufio"
	"cmp"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

type sortEntry struct {
	OriginalValue string
	DateValue     *time.Time
	FloatValue    float64
}

var timeVal = regexp.MustCompile("(?i)(pm|15:)")

func (e *EditorArgs) Sort(input string) (string, error) {
	nums, dates, strs := e.getSortEntries(input)

	slices.SortFunc(nums, func(a, b sortEntry) int {
		if e.Switches.Invert {
			return cmp.Compare(b.FloatValue, a.FloatValue)
		}
		return cmp.Compare(a.FloatValue, b.FloatValue)
	})

	slices.SortFunc(strs, func(a, b string) int {
		if e.Option == "i" {
			if e.Switches.Invert {
				return strings.Compare(strings.ToLower(b), strings.ToLower(a))
			}
			return strings.Compare(strings.ToLower(a), strings.ToLower(b))
		}

		if e.Switches.Invert {
			return strings.Compare(b, a)
		}

		return strings.Compare(a, b)
	})

	slices.SortFunc(dates, func(a, b sortEntry) int {
		if e.Switches.Invert {
			return cmp.Compare(b.DateValue.UnixMicro(), a.DateValue.UnixMicro())
		}
		return cmp.Compare(a.DateValue.UnixMicro(), b.DateValue.UnixMicro())
	})

	output := make([]string, 0)
	strdates := make([]string, len(dates))
	for d := range dates {
		strdates[d] = dates[d].OriginalValue
	}
	strnums := make([]string, len(nums))
	for n := range nums {
		strnums[n] = nums[n].OriginalValue
	}

	if len(strnums) > 0 {
		output = append(output, strings.Join(strnums, e.RowDelimiter))
	}

	if len(strdates) > 0 {
		output = append(output, strings.Join(strdates, e.RowDelimiter))
	}

	if len(strs) > 0 {
		output = append(output, strings.Join(strs, e.RowDelimiter))
	}

	return strings.Join(output, e.RowDelimiter), nil
}

func (e *EditorArgs) bufSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), e.RowDelimiter); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}


func (e *EditorArgs) getSortEntries(input string) ([]sortEntry, []sortEntry, []string){
	lineScanner := bufio.NewScanner(strings.NewReader(input))
	lineScanner.Split(e.bufSplit)
	nums := make([]sortEntry, 0)
	times := make([]sortEntry, 0)
	strings := make([]string, 0)
	
	for lineScanner.Scan() {
		originalValue := lineScanner.Text()

		date := getDateValue(originalValue)
		if date != nil {
			times = append(times, sortEntry{ DateValue: date, OriginalValue: originalValue })
			continue
		}

		flt, err := strconv.ParseFloat(originalValue, 64)
		if err == nil {
			nums = append(nums, sortEntry{ FloatValue: flt, OriginalValue: originalValue })
			continue
		}

		strings = append(strings, originalValue)
	}

	return nums, times, strings
}

func getDateValue(input string) *time.Time {
	for _, dfmt := range settings().DateFormats {
		date, err := time.Parse(dfmt, input)
		if err == nil {
			return &date
		}

		//if it contains time information, don't append time
		if timeVal.MatchString(dfmt) {
			continue
		}

		for _, tfmt := range settings().TimeFormats {
			dtfmt := fmt.Sprintf("%s %s", dfmt, tfmt)
			date, err = time.Parse(dtfmt, input)
			if err == nil {
				return &date
			}

			dtfmt = fmt.Sprintf("%sT%s", dfmt, tfmt)
			date, err = time.Parse(dtfmt, input)
			if err == nil {
				return &date
			}
		}
	}

	for _, tfmt := range settings().TimeFormats {
		date, err := time.Parse(tfmt, input)
		if err == nil {
			return &date
		}
	}

	return nil
}
