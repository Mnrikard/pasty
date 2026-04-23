package edit

import (
	"fmt"

	"github.com/spf13/cobra"
)

type SubCommand struct {
	Name               string
	Use                string
	Short              string
	Long               string
	Aliases            []string
	Args               cobra.PositionalArgs
	ArgDefs            []Arg
	EditFunc           func(*EditorArgs) func(string) (string, error)
	NeedsRegexSwitches bool
	CustomSetup        func(*cobra.Command, *EditorArgs)
}

var SubCommands = []SubCommand{
	{
		Name:      "plugin",
		Use:       "plugin [name] [args...]",
		Short:     "Runs a plugin found in your config directory",
		Long:      "Define plugins in $HOME/.config/pasty/plugins/{name}.lua to be executed here",
		Args:      cobra.MinimumNArgs(1),
		ArgDefs: []Arg{
			{
				Position:     0,
				HelpText:     "Name of Plugin",
				SetValue: func(e *EditorArgs, value string) {
					e.Option = value
				},
			},
		},
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.HandlePlugin },
	},
	{
		Name:      "udf",
		Use:       "udf [name] [args...]",
		Short:     "Runs a user defined function",
		Args:      cobra.MinimumNArgs(1),
		ArgDefs: []Arg{
			{
				Position:     0,
				HelpText:     "Name of UDF",
				SetValue: func(e *EditorArgs, value string) {
					e.Option = value
				},
			},
		},
	},
	{
		Name:      "upper",
		Use:      "upper",
		Short:    "Upper cases text",
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.Upper },
	},
	{
		Name:     "lower",
		Use:      "lower",
		Short:    "Lower cases text",
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.Lower },
	},
	{
		Name:     "title",
		Use:      "title",
		Short:    "Title cases text",
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.Title },
	},
	{
		Name:     "base64encode",
		Use:      "base64encode",
		Aliases:  []string{"base64"},
		Short:    "Base64 encodes the text",
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.EncodeBase64 },
	},
	{
		Name:     "base64decode",
		Use:      "base64decode",
		Short:    "Base64 decodes the text",
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.DecodeBase64 },
	},
	{
		Name:     "urlencode",
		Use:      "urlencode",
		Aliases:  []string{"url"},
		Short:    "Url encodes the text",
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.EncodeForUrl },
	},
	{
		Name:     "urldecode",
		Use:      "urldecode",
		Short:    "Url decodes the text",
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.DecodeFromUrl },
	},
	{
		Name:     "xmlencode",
		Use:      "xmlencode",
		Short:    "XML encodes the text",
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.EncodeForXml },
	},
	{
		Name:     "xmldecode",
		Use:      "xmldecode",
		Short:    "XML decodes the text",
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.DecodeFromXml },
	},
	{
		Name:     "math",
		Use:      "math",
		Short:    "Evaluates simple math equations",
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.SolveMath },
	},
	{
		Name:     "newGuid",
		Use:      "newGuid",
		Aliases:  []string{"newid"},
		Short:    "Creates a new v4 GUID",
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.NewGuid },
	},
	{
		Name:  "columnAlign",
		Use:   "columnAlign",
		Short: "Aligns columns",
		Args:  cobra.MinimumNArgs(0),
		ArgDefs: []Arg{
			{
				Position:     0,
				HelpText:     "Number of Spaces Between Columns",
				DefaultValue: "2",
				SetValue: func(e *EditorArgs, value string) {
					var err error
					e.NumSpaces, err = parseIntArg(value)
					if err != nil {
						displayArgError(err)
					}
				},
			},
			{
				Position:     1,
				HelpText:     "Delimiter (defaults to tab)",
				DefaultValue: "\t",
				SetValue: func(e *EditorArgs, value string) {
					e.ColumnDelimiter = value
				},
			},
		},
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.AlignColumns },
	},
	{
		Name:    "count",
		Use:     "count",
		Aliases: []string{"length", "len"},
		Short:   "Gets the count of characters, lines, or words",
		Args:    cobra.MinimumNArgs(0),
		ArgDefs: []Arg{
			{
				Position:     0,
				Options:      []string{"lines", "words", "chars"},
				DefaultValue: "characters",
				SetValue: func(e *EditorArgs, value string) {
					e.Option = value
				},
			},
		},
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.CountItem },
	},
	{
		Name:  "dedup",
		Use:   "dedup [separator]",
		Short: "Deduplicates items",
		Long: `Deduplicates items
example:
$> echo "hello hello world" | pasty dedup " "
hello world
`,
		Args: cobra.MinimumNArgs(0),
		ArgDefs: []Arg{
			{
				HelpText:     "Separator (default new line)",
				DefaultValue: "\n",
				SetValue: func(e *EditorArgs, s string) {
					e.RowDelimiter = s
				},
			},
		},
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.Deduplicate },
	},
	{
		Name:  "insert",
		Use:   "insert",
		Short: "Converts result sets into an insert statement",
		Args:  cobra.MinimumNArgs(2),
		ArgDefs: []Arg{
			{
				Position: 0,
				HelpText: "Table Name",
				SetValue: func(e *EditorArgs, value string) {
					e.Option = value
				},
			},
			{
				Position: 1,
				HelpText: "Column Delimiter",
				SetValue: func(e *EditorArgs, value string) {
					e.ColumnDelimiter = value
				},
			},
			{
				Position:     2,
				HelpText:     "Row Delimiter",
				DefaultValue: "\r?\n",
				SetValue: func(e *EditorArgs, value string) {
					e.RowDelimiter = value
				},
			},
		},
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.InsertSQL },
	},
	{
		Name:  "setText",
		Use:   "setText",
		Short: "Sets the text to the given string",
		Args:  cobra.MinimumNArgs(1),
		ArgDefs: []Arg{
			{
				Position: 0,
				HelpText: "Text to set",
				SetValue: func(e *EditorArgs, value string) {
					e.Option = value
				},
			},
		},
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.SetText },
	},
	{
		Name:  "toBase",
		Use:   "toBase",
		Short: "converts base 10 to the given base",
		Args:  cobra.MinimumNArgs(1),
		ArgDefs: []Arg{
			{
				Position: 0,
				HelpText: "Base",
				SetValue: func(e *EditorArgs, value string) {
					e.Option = value
				},
			},
		},
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.ToNumBase },
	},
	{
		Name:  "fromBase",
		Use:   "fromBase",
		Short: "converts from the given base to base 10",
		Args:  cobra.MinimumNArgs(1),
		ArgDefs: []Arg{
			{
				Position: 0,
				HelpText: "Base",
				SetValue: func(e *EditorArgs, value string) {
					e.Option = value
				},
			},
		},
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.FromNumBase },
	},
	{
		Name:               "rep",
		Use:                "rep <regex match> <replacement string> [regex switches]",
		Short:              "replace text",
		Args:               cobra.MinimumNArgs(2),
		NeedsRegexSwitches: true,
		ArgDefs: []Arg{
			{
				Position: 0,
				HelpText: "Regex Pattern to Replace",
				SetValue: func(e *EditorArgs, value string) {
					e.Regex = value
				},
			},
			{
				Position: 1,
				HelpText: "Replacement Pattern",
				SetValue: func(e *EditorArgs, value string) {
					e.Replacement = value
				},
			},
		},
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.ReplaceText },
	},
	{
		Name:  "grep",
		Use:   "grep <regex match> [separator] [regex switches]",
		Short: "Finds matches from the input text",
		Long: `By default, tests each line of input and returns that line if it matches the given regular expression
* If the "o" switch is provided, returns the matches only, not the matched lines
* If the "v" switch is provided, inverts the matches
`,
		Args:               cobra.MinimumNArgs(1),
		NeedsRegexSwitches: true,
		ArgDefs: []Arg{
			{
				Position: 0,
				HelpText: "Regex Pattern to Find",
				SetValue: func(e *EditorArgs, value string) {
					e.Regex = value
				},
			},
			{
				Position:     1,
				HelpText:     "Separator Text (defaults to new line)",
				DefaultValue: "\n",
				SetValue: func(e *EditorArgs, value string) {
					e.RowDelimiter = value
				},
			},
		},
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.Grep },
	},
	{
		Name:  "sort",
		Use:   "sort",
		Short: "Sorts alphabetically and numerically",
		Long: `Sorts sets alphabetically and numerically`,
		Args:               cobra.MinimumNArgs(0),
		ArgDefs: []Arg{
			{
				Position: 0,
				HelpText: "Separator",
				SetValue: func(e *EditorArgs, value string) {
					e.RowDelimiter = value
				},
				DefaultValue: "\n",
			},
		},
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.Sort },
	},
}

func parseIntArg(value string) (int, error) {
	var result int
	_, err := fmt.Sscanf(value, "%d", &result)
	return result, err
}

func displayArgError(err error) {
	fmt.Println(err)
}
