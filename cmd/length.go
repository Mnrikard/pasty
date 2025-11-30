package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mattr/pasty/text"
	"github.com/mattr/pasty/util"
	"github.com/spf13/cobra"
	"github.com/gen2brain/beeep"
)


var lengthArgDefs = []util.Arg {
	{
		Position: 0,
		Options: []string { "lines", "words", "chars" },
		SetValue: func(e *util.Editor, value string) {
			e.Option = value
		},
		DefaultValue: "characters",
	},
}

var Length = &util.Editor {
	ArgDefs: lengthArgDefs,
	Command: &cobra.Command {
		ValidArgsFunction: util.BuildArguments(lengthArgDefs),
		Run: func(cmd *cobra.Command, args []string) {
			e := util.Editor{}
			util.GetArguments(&e, lengthArgDefs, args)
			countItem(e)
		},
		Use:   "length",
		Short: "Gets the length of characters, lines, or words",
		Args:  cobra.MinimumNArgs(0),
	},
}

func countItem(e util.Editor) {
	txt, err := text.GetText();
	if err != nil {
		panic(err)
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

	beeep.AppName = "pasty"
	var icon []byte
	err = beeep.Alert(e.Option, fmt.Sprintf("%d %s\n", count, e.Option), icon)
	if err != nil {
		panic(err)
	}
}
