package cmd

import (
	"github.com/mattr/pasty/edit"
	"github.com/mattr/pasty/text"
	"github.com/spf13/cobra"
)

var countArgDefs = []edit.Arg {
	{
		Position: 0,
		Options: []string { "lines", "words", "chars" },
		SetValue: func(e *edit.EditorArgs, value string) {
			e.Option = value
		},
		DefaultValue: "characters",
	},
}

var Counter = &cobra.Command {
	ValidArgsFunction: edit.BuildArguments(countArgDefs),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		e.GetArguments(countArgDefs, args)
		text.EditText(e, e.CountItem)
	},
	Use:   "count",
	Aliases: []string{"length","len"},
	Short: "Gets the count of characters, lines, or words",
	Args:  cobra.MinimumNArgs(0),
}

