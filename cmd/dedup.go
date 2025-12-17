package cmd

import (
	"github.com/mattr/pasty/edit"
	"github.com/mattr/pasty/text"
	"github.com/spf13/cobra"
)

var dedupArgs = []edit.Arg {
	{
		HelpText: "Separator (default new line)",
		SetValue: func(e *edit.EditorArgs, s string) {
			e.RowDelimiter = s
		},
		DefaultValue: "\n",
	},
}

var Dedup = &cobra.Command{
	Use: "dedup [separator]",
	Short: "Deduplicates items",
	Long: `Deduplicates items
example:
$> echo "hello hello world" | pasty dedup " "
hello world
`,
	Args: cobra.MinimumNArgs(0),
	ValidArgsFunction: edit.BuildArguments(dedupArgs),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		e.GetArguments(dedupArgs, args)
		text.EditText(e, e.Deduplicate)
	},
}
