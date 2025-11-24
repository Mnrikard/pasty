package cmd

import (
	"fmt"
	"regexp"

	"github.com/mattr/pasty/text"
	"github.com/mattr/pasty/util"
	"github.com/spf13/cobra"
)

var argfunc = util.BuildArguments([]util.Arg{
	util.Arg { Position: 0, HelpText: "Regex Pattern to Replace" },
	util.Arg { Position: 1, HelpText: "Replacement Pattern" },
	util.Arg { Position: 2, Options: []string{"i", "m", "s", "U"} },
})

var repCmd = &cobra.Command{
	Use:   "rep <regex match> <replacement string> [regex switches]",
	Short: "replace text",
	Args:  cobra.MinimumNArgs(2),
	ValidArgsFunction: argfunc,
	Run: func(cmd *cobra.Command, args []string) {
		replacement := ""
		if len(args) > 1 {
			replacement = args[1]
		}
		regex := args[0]
		if len(args) > 2 {
			regex = "(?" + args[2] + ")" + regex
		}
		replaceText(regex, replacement)
	},
}

func replaceText(regex string, repla string) {
	txt, err := text.GetText();
	if err != nil {
		fmt.Println(err)
	}

	rx, err := regexp.Compile(regex)
	if err != nil {
		panic(err)
	}

	replacedText := rx.ReplaceAllString(txt, repla)
	err = text.SetText(replacedText)
	if err != nil {
		panic(err)
	}
}
