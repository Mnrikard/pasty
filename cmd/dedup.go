package cmd

import (
	"slices"
	"strings"

	"github.com/mattr/pasty/text"
	"github.com/mattr/pasty/util"
	"github.com/spf13/cobra"
)

var dedupArgs = []util.Arg {
	{
		HelpText: "Separator (default new line)",
		SetValue: func(e *util.Editor, s string) {
			e.RowDelimiter = s
		},
		DefaultValue: "\n",
	},
}

var Dedup = &util.Editor{
	ArgDefs: dedupArgs,
	Command: &cobra.Command{
		Use: "dedup [separator]",
		Short: "Deduplicates items",
		Long: `Deduplicates items
example:
$> echo "hello hello world" | pasty dedup " "
hello world
`,
		Args: cobra.MinimumNArgs(0),
		ValidArgsFunction: util.BuildArguments(dedupArgs),
		Run: func(cmd *cobra.Command, args []string) {
			e := &util.Editor{}
			util.GetArguments(e, dedupArgs, args)
			deduplicate(e)
		},
	},
}

func deduplicate(e *util.Editor) {
	txt, err := text.GetText();
	if err != nil {
		util.DisplayError(err)
	}

	items := strings.Split(txt, e.RowDelimiter)
	newItems := make([]string, 0)
	for _, item := range items {
		if !slices.Contains(newItems, item) {
			newItems = append(newItems, item)
		}
	}

	replacedText := strings.Join(newItems, e.RowDelimiter)

	err = text.SetText(replacedText)
	if err != nil {
		util.DisplayError(err)
	}
}
