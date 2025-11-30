package cmd

import (
	"strconv"
	"strings"

	"github.com/mattr/pasty/text"
	"github.com/mattr/pasty/util"
	"github.com/spf13/cobra"
)

var columnAlignArgs = []util.Arg {
	{
		Position: 0,
		HelpText: "Number of Spaces Between Columns",
		SetValue: func(e *util.Editor, value string) {
			var err error
			e.NumSpaces, err = strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
		},
		DefaultValue: "2",
	},
	{
		Position: 1,
		HelpText: "Delimiter (defaults to tab)",
		SetValue: func(e *util.Editor, value string) {
			e.ColumnDelimiter = value
		},
		DefaultValue: "\t",
	},
}

var ColumnAlign = &util.Editor {
	ArgDefs: columnAlignArgs,
	Command: &cobra.Command {
		ValidArgsFunction: util.BuildArguments(columnAlignArgs),
		Run: func(cmd *cobra.Command, args []string) {
			e := util.Editor{}
			util.GetArguments(&e, columnAlignArgs, args)
			alignColumns(e)
		},
		Use:   "columnAlign",
		Short: "Aligns columns",
		Args:  cobra.MinimumNArgs(0),
	},
}

func alignColumns(e util.Editor) {
	txt, err := text.GetText();
	if err != nil {
		panic(err)
	}

	rows := util.SplitRows(txt)
	grid := util.SplitColumns(rows, e.ColumnDelimiter)
	var colWidths = getColWidths(grid)
	replacedText := rebuildRows(e, grid, colWidths)


	err = text.SetText(replacedText)
	if err != nil {
		panic(err)
	}
}

func rebuildRows(e util.Editor, grid [][]string, colWidths []int) string {
	var output = make([]string, len(grid))

	for ri, row := range grid {
		cols := make([]string, 0)
		for ci, col := range row {
			cols = append(cols, col)
			if len(col) < colWidths[ci] {
				cols = append(cols, strings.Repeat(" ", colWidths[ci]-len(col)))
			}
			cols = append(cols, strings.Repeat(" ", e.NumSpaces))
		}

		output[ri] = strings.Join(cols, "")
	}

	//todo: global settings for new line
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
