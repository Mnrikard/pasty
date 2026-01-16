package cmd

import (
	"github.com/mattr/pasty/edit"
	"github.com/mattr/pasty/text"
	"github.com/spf13/cobra"
)

var setTextArgDefs = []edit.Arg {
	{
		Position: 0,
		HelpText: "Text to set",
		SetValue: func(e *edit.EditorArgs, value string) {
			e.Option = value
		},
	},
}

var SetText = &cobra.Command {
	ValidArgsFunction: edit.BuildArguments(setTextArgDefs),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		e.GetArguments(setTextArgDefs, args)
		text.EditText(e, e.SetText)
	},
	Use:   "setText",
	Short: "Sets the text to the given string",
	Args:  cobra.MinimumNArgs(1),
}

var NewGuid = &cobra.Command {
	Use:   "newGuid",
	Aliases: []string{"newid"},
	Short: "Creates a new v4 GUID",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		text.EditText(e, e.NewGuid)
	},
}

