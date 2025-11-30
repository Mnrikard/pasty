package util

import (
	"log"
	"regexp"
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
