package cmd

import (
	"strconv"

	"github.com/mattr/pasty/edit"
	"github.com/mattr/pasty/text"
	"github.com/mattr/pasty/util"
	"github.com/spf13/cobra"
)

var columnAlignArgs = []edit.Arg {
	{
		Position: 0,
		HelpText: "Number of Spaces Between Columns",
		SetValue: func(e *edit.EditorArgs, value string) {
			var err error
			e.NumSpaces, err = strconv.Atoi(value)
			if err != nil {
				util.DisplayError(err)
			}
		},
		DefaultValue: "2",
	},
	{
		Position: 1,
		HelpText: "Delimiter (defaults to tab)",
		SetValue: func(e *edit.EditorArgs, value string) {
			e.ColumnDelimiter = value
		},
		DefaultValue: "\t",
	},
}

var ColumnAlign = &cobra.Command {
	ValidArgsFunction: edit.BuildArguments(columnAlignArgs),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		e.GetArguments(columnAlignArgs, args)
		text.EditText(e, e.AlignColumns)
	},
	Use:   "columnAlign",
	Short: "Aligns columns",
	Args:  cobra.MinimumNArgs(0),
}

