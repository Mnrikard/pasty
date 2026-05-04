package edit

import (
	"strings"

	"github.com/mattr/pasty/util"
)

func (e *EditorArgs) AlignColumns(input string) (string, error) {
	rows := util.SplitRows(input)
	grid := util.SplitColumns(rows, e.ColumnDelimiter)
	var colWidths = getColWidths(grid)
	replacedText := rebuildRows(e, grid, colWidths)
	return replacedText, nil
}

func rebuildRows(e *EditorArgs, grid [][]string, colWidths []int) string {
	var output = make([]string, len(grid))

	for ri, row := range grid {
		cols := make([]string, 0)
		for ci, col := range row {
			cols = append(cols, col)
			if len(col) < colWidths[ci] {
				cols = append(cols, strings.Repeat(" ", colWidths[ci]-len(col)))
			}

			if ci < len(row)-1 {
				cols = append(cols, strings.Repeat(" ", e.NumSpaces))
			}
		}

		output[ri] = strings.Join(cols, "")
	}

	return strings.Join(output, "\n")
}

func getColWidths(grid [][]string) []int {
	if len(grid) < 1 {
		return nil
	}

	widths := make([]int, len(grid[0]))
	for _, row := range grid {
		for ci, col := range row {
			if ci+1 > len(widths) {
				widths = append(widths, len(col))
			}
			if len(col) > widths[ci] {
				widths[ci] = len(col)
			}
		}
	}

	return widths
}
