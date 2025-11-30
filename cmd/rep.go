package cmd

import (
	"regexp"

	"github.com/mattr/pasty/text"
	"github.com/mattr/pasty/util"
	"github.com/spf13/cobra"
)

var repArgDefs = []util.Arg{
	{
		Position: 0,
		HelpText: "Regex Pattern to Replace",
		SetValue: func (e *util.Editor, value string) {
			e.Regex = value
		},
	},
	{
		Position: 1,
		HelpText: "Replacement Pattern",
		SetValue: func (e *util.Editor, value string) {
			e.Replacement = value
		},
	},
	{
		Position: 2,
		Options: []string{"i", "m", "s", "U"},
		SetValue: func (e *util.Editor, value string) {
			e.Regex = "(?" + value + ")" + e.Regex
		},
	},
}

var Replace = &util.Editor {
	ArgDefs: repArgDefs,
	Command: &cobra.Command{
		Use:   "rep <regex match> <replacement string> [regex switches]",
		Short: "replace text",
		Args:  cobra.MinimumNArgs(2),
		ValidArgsFunction: util.BuildArguments(repArgDefs),
		Run: func(cmd *cobra.Command, args []string) {
			e := &util.Editor{}
			util.GetArguments(e, repArgDefs, args)
			editorFuncRep(e.Regex, e.Replacement)
		},
	},
}

func editorFuncRep(regex string, repla string) {
	txt, err := text.GetText();
	if err != nil {
		util.DisplayError(err)
	}

	rx, err := regexp.Compile(regex)
	if err != nil {
		util.DisplayError(err)
	}

	replacedText := rx.ReplaceAllString(txt, repla)
	err = text.SetText(replacedText)
	if err != nil {
		util.DisplayError(err)
	}
}
