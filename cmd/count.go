package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mattr/pasty/text"
	"github.com/mattr/pasty/util"
	"github.com/spf13/cobra"
)

var countArgDefs = []util.Arg {
	{
		Position: 0,
		Options: []string { "lines", "words", "chars" },
		SetValue: func(e *util.Editor, value string) {
			e.Option = value
		},
		DefaultValue: "characters",
	},
}

var Counter = &util.Editor {
	ArgDefs: countArgDefs,
	Command: &cobra.Command {
		ValidArgsFunction: util.BuildArguments(countArgDefs),
		Run: func(cmd *cobra.Command, args []string) {
			e := util.Editor{}
			util.GetArguments(&e, countArgDefs, args)
			countItem(e)
		},
		Use:   "count",
		Aliases: []string{"length","len"},
		Short: "Gets the count of characters, lines, or words",
		Args:  cobra.MinimumNArgs(0),
	},
}

func countItem(e util.Editor) {
	txt, err := text.GetText();
	if err != nil {
		util.DisplayError(err)
	}

	var count int
	var rx *regexp.Regexp

	switch e.Option {
	case "words","word":
		rx = regexp.MustCompile(`\s+`)
		count = len(rx.Split(strings.TrimSpace(txt), -1))
	case "lines","line":
		rx = regexp.MustCompile("\r?\n")
		count = len(rx.Split(strings.TrimSpace(txt), -1))
	default:
		count = len(strings.TrimSpace(txt))
	}

	util.Notify(e.Option, fmt.Sprintf("%d %s\n", count, e.Option))
}
