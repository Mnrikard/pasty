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
* If the "o" switch is provided, returns the matches only, not the matched lines
* If the "v" switch is provided, inverts the matches
`,
	Args:  cobra.MinimumNArgs(1),
	ValidArgsFunction: edit.BuildArguments(grepArgDefs),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		e.GetArguments(grepArgDefs, args)
		if rootSwitches.GrepOnlyMatching {
			e.Option = "OnlyMatching"
		}
		if rootSwitches.GrepInvertMatch {
			e.Invert = true
		}
		text.EditText(e, e.Grep)
	},
}


