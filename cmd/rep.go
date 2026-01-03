package cmd

import (
	"github.com/mattr/pasty/edit"
	"github.com/mattr/pasty/text"
	"github.com/spf13/cobra"
)

var repArgDefs = []edit.Arg{
	{
		Position: 0,
		HelpText: "Regex Pattern to Replace",
		SetValue: func (e *edit.EditorArgs, value string) {
			e.Regex = value
		},
	},
	{
		Position: 1,
		HelpText: "Replacement Pattern",
		SetValue: func (e *edit.EditorArgs, value string) {
			e.Replacement = value
		},
	},
}

var Replace = &cobra.Command{
	Use:   "rep <regex match> <replacement string> [regex switches]",
	Short: "replace text",
	Args:  cobra.MinimumNArgs(2),
	ValidArgsFunction: edit.BuildArguments(repArgDefs),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		e.GetArguments(repArgDefs, args)
		e.PrependRegex(rootSwitches)
		text.EditText(e, e.ReplaceText)
	},
}
