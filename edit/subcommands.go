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
		Name:  "plugin",
		Use:   "plugin [name] [args...]",
		Short: "Runs a plugin found in ~/.config/pasty/plugins/{name}.lua",
		Long:  "Define plugins in $HOME/.config/pasty/plugins/{name}.lua to be executed here",
		Args:  cobra.MinimumNArgs(1),
		ArgDefs: []Arg{
			{
				Position: 0,
				HelpText: "Name of Plugin",
				SetValue: func(e *EditorArgs, value string) {
					e.Option = value
				},
			},
		},
	},
	{
		Name:  "udf",
		Use:   "udf [name] [args...]",
		Short: "Runs a user defined function found in ~/.config/pasty/settings.json",
		Long:  "See documentation for structuring these functions",
		Args:  cobra.MinimumNArgs(1),
		ArgDefs: []Arg{
			{
				Position: 0,
				HelpText: "Name of UDF",
				SetValue: func(e *EditorArgs, value string) {
					e.Option = value
				},
			},
		},
	},
	{
		Name:  "upper",
		Use:   "upper",
		Short: "Capitalizes all text",
		Long: `Syntax: pasty upper

	Example: echo abcd | pasty upper
	>> ABCD
`,
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.Upper },
	},
	{
		Name:  "lower",
		Use:   "lower",
		Short: "Lower cases all text",
		Long: `Syntax: pasty lower

	Example: echo ABCD | pasty lower
	>> abcd
`,
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.Lower },
	},
	{
		Name:  "title",
		Use:   "title",
		Short: "Title cases text",
		Long: `Syntax: pasty title

	Example: echo "the quick brown fox jumped over the lazy dog" | pasty lower
	>> The Quick Brown Fox Jumped Over The Lazy Dog
`,
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.Title },
	},
	{
		Name:    "base64encode",
		Use:     "base64encode",
		Aliases: []string{"base64"},
		Short:   "Base64 encodes the text",
		Long: `Syntax: pasty base64encode

	Example: echo "username:password" | pasty base64encode
	>> dXNlcm5hbWU6cGFzc3dvcmQK
`,
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.EncodeBase64 },
	},
	{
		Name:  "base64decode",
		Use:   "base64decode",
		Short: "Decodes the text from base64 if possible",
		Long: `Syntax: pasty base64decode

	Example: echo "dXNlcm5hbWU6cGFzc3dvcmQK" | pasty base64decode
	>> username:password
`,
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.DecodeBase64 },
	},
	{
		Name:    "urlencode",
		Use:     "urlencode",
		Aliases: []string{"url"},
		Short:   "Url encodes the text",
		Long: `Syntax: pasty urlencode

	Example: echo "this&that" | pasty urlencode
	>> this%26that
`,
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.EncodeForUrl },
	},
	{
		Name:  "urldecode",
		Use:   "urldecode",
		Short: "Url decodes the text",
		Long: `Syntax: pasty urldecode

	Example: echo "this%26that%2F0" | pasty urldecode
	>> this&that/0
`,
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.DecodeFromUrl },
	},
	{
		Name:  "xmlencode",
		Use:   "xmlencode",
		Short: "XML encodes the text",
		Long: `Syntax: pasty xmlencode

	Example: echo "this > that & other things" | pasty xmlencode
	>> this &gt; that &amp; other things
`,
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.EncodeForXml },
	},
	{
		Name:  "xmldecode",
		Use:   "xmldecode",
		Short: "XML decodes the text",
		Long: `Syntax: pasty xmldecode

	Example: echo "this &gt; that &amp; other things" | pasty xmldecode
	>> this > that & other things
`,
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.DecodeFromXml },
	},
	{
		Name:  "jwtdecode",
		Use:   "jwtdecode",
		Short: "Decodes a JWT token",
		Long: `Syntax: pasty jwtdecode [key]

	Passing in a key will validate the signature and the expiration
`,
		Args:     cobra.MinimumNArgs(0),
		ArgDefs: []Arg{
			{
				Position:     0,
				HelpText:     "Key",
				DefaultValue: "",
				SetValue: func(e *EditorArgs, value string) {
					e.Option = value
				},
			},
		},
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.JwtDecode },
	},
	{
		Name:  "math",
		Use:   "math",
		Short: "Evaluates simple math equations",
		Long: `Syntax: pasty math

	Example: echo "1+(1+2)/3" | pasty math
	>> 2
`,
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.SolveMath },
	},
	{
		Name:    "newGuid",
		Use:     "newGuid",
		Aliases: []string{"newid"},
		Short:   "Creates a new v4 GUID",
		Long: `Syntax: pasty <newGuid|newid>

	Places a new v4 GUID on your clipboard or standard output
`,
		Args:     cobra.MinimumNArgs(0),
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.NewGuid },
	},
	{
		Name:  "columnAlign",
		Use:   "columnAlign",
		Short: "Aligns delimited data by columns",
		Long: `Syntax: pasty columnAlign ["number of spaces between rows"] ["input delimiter"]
	Number of spaces defaults to 2
	input delimiter defaults to tab character

	Example: cat tabDelimited.file | pasty columnAlign 2 "\t"
	>> col1   col2             col3
	>> names  some other data  1234
`,
		Args: cobra.MinimumNArgs(0),
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
		Long: `Syntax: pasty <count|len|length> [lines|words|chars]
	defaults to counting characters

	Example: echo abcd | pasty len
	>> 4 characters
`,
		Args: cobra.MinimumNArgs(0),
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
		Long: `Syntax: pasty dedup [separator]
	Separator defaults to a new line

	Example: echo "hello hello world" | pasty dedup " "
	>> hello world
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
		Long: `Syntax: pasty insert mydb.dbo.mytable [delimiter]
	delimiter defaults to tab character
	SQL server specific, groups rows into collections of 1000 inserts at a time

	Example: pasty insert mydb.dbo.mytable
	>> insert into [mydb].[dbo].[mytable] (col1, col2) values ('val1','val2');
`,
		Args: cobra.MinimumNArgs(2),
		ArgDefs: []Arg{
			{
				Position: 0,
				HelpText: "Table Name",
				SetValue: func(e *EditorArgs, value string) {
					e.Option = value
				},
			},
			{
				Position:     1,
				HelpText:     "Column Delimiter",
				DefaultValue: "\t",
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
		Name:    "format",
		Use:     "format [type]",
		Aliases: []string{"fmt"},
		Short:   "Performs a simple format on specific data or code",
		Long: `Syntax: pasty format [json|sql]
	Performs a simple, opinionated format of data
	There are other, better solutions for formatting code that are geared towards that. This is just a quick and dirty format for ease of use

	Example: echo "{'name':'value'}" | pasty format json
	>> {
	>>   'name':'value;
	>> }
`,
		Args: cobra.MinimumNArgs(0),
		ArgDefs: []Arg{
			{
				Position: 0,
				Options:  []string{"sql", "json"},
				SetValue: func(e *EditorArgs, value string) {
					e.Option = value
				},
			},
		},
		EditFunc: func(e *EditorArgs) func(string) (string, error) { return e.FormatCode },
	},
	{
		Name:  "setText",
		Use:   "setText",
		Short: "Sets the text to the given string",
		Long: `Syntax: pasty setText [some text]
	Useful for building user defined functions (see pasty udf)
	by assigning a particular value to the input at the beginning

	Example: echo "anything at all" | pasty setText "something else"
	>> something else
`,
		Args: cobra.MinimumNArgs(1),
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
		Long: `Syntax: pasty tobase [integer base]

	Example: echo "16" | pasty tobase 16
	>> F
`,
		Args: cobra.MinimumNArgs(1),
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
		Long: `Syntax: pasty fromBase [integer base]

	Example: echo "F" | pasty fromBase 16
	>> 16
`,
		Args: cobra.MinimumNArgs(1),
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
		Name:  "rep",
		Use:   "rep <regex match> <replacement string> [regex switches]",
		Short: "Replaces text with a regular expression",
		Long: `Syntax: pasty rep <pattern> <replacement>

	Example: echo "sw33t" | pasty rep "\\d" "e"
	>> sweet
`,
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

	Syntax: pasty grep "pattern"

	Example: echo "1. Value\n2. Text\nA. sublist" | pasty grep "\\d"
	>> 1. Value
	>> 2. Text
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
		Long: `Sorts sets alphabetically and numerically
	Syntax: pasty sort [separator]
	separator defaults to new line

	Example: echo "1,3,2,4,6,5" | pasty sort ","
	>> 1,2,3,4,5,6
`,
		Args: cobra.MinimumNArgs(0),
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
