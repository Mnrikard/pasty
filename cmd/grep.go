package cmd

import (
	"github.com/mattr/pasty/edit"
	"github.com/mattr/pasty/text"
	"github.com/spf13/cobra"
)

var grepArgDefs = []edit.Arg{
	{
		Position: 0,
		HelpText: "Regex Pattern to Find",
		SetValue: func (e *edit.EditorArgs, value string) {
			e.Regex = value
		},
	},
	{
		Position: 1,
		HelpText: "Separator Text (defaults to new line)",
		SetValue: func (e *edit.EditorArgs, value string) {
			e.RowDelimiter = value
		},
		DefaultValue: "\n",
	},
}

var Grep = &cobra.Command{
	Use:   "grep <regex match> [separator] [regex switches]",
	Short: "Finds matches from the input text",
	Long: `By default, tests each line of input and returns that line if it matches the given regular expression
* If the "L" switch is provided, returns the matchs only, not the matched lines
`,
	Args:  cobra.MinimumNArgs(1),
	ValidArgsFunction: edit.BuildArguments(repArgDefs),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		e.GetArguments(repArgDefs, args)
		if rootSwitches.GrepAll {
			e.Option = "L"
		}
		text.EditText(e, e.Grep)
	},
}


