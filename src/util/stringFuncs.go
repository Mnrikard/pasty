package util

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func SplitRows(input string) []string {
	rx, _ := regexp.Compile("\r?\n")
	log.Printf("%v", rx)
	return rx.Split(input, -1)
}

func SplitColumns(rows []string, columnDelimiter string) [][]string {
	var output [][]string
	output = make([][]string, len(rows))
	for ri, row := range rows {
		output[ri] = strings.Split(row, columnDelimiter)
	}

	return output
}

func IsNullOrNumber(input string) bool {
	if len(input) == 0 {
		return false
	}

	if strings.ToUpper(input) == "NULL" {
		return true
	}

	_, numErr := strconv.ParseFloat(input, 64)
	return numErr == nil
}

func EscapeRegex(input string) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(
						strings.ReplaceAll(
							strings.ReplaceAll(input,
								"(", "\\("),
							")", "\\)"),
						"+", "\\+"),
					"*", "\\*"),
				"-", "\\-"),
			".", "\\."),
		"|", "\\|")
}
