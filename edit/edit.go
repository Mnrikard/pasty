package edit

import (
	"regexp"
	"strings"

	"github.com/Mnrikard/pasty/switches"
	"github.com/spf13/cobra"
)

var existingFlagsRx = regexp.MustCompile("^\\(\\?([sUmi]+)(-([sUmi]+))?\\)")

type Arg struct {
	Position     int
	HelpText     string
	Options      []string
	GetOptions   func() []string
	SetValue     func(*EditorArgs, string)
	DefaultValue string
}

type EditorArgs struct {
	Regex           string
	Replacement     string
	ColumnDelimiter string
	RowDelimiter    string
	NumSpaces       int
	Option          string
	Key             string
	Invert          bool
	OriginalArgs    []string
	Switches        *switches.Switches
	Regexes         map[string]*regexp.Regexp
}

func (e *EditorArgs) PrependRegex() {
	readExistingFlags(e)
	sw := make([]string, 0)
	if e.Switches == nil {
		return
	}
	if !e.Switches.CaseSensitive {
		sw = append(sw, "i")
	}
	if e.Switches.SingleLine {
		sw = append(sw, "s")
	}
	if e.Switches.MultiLine {
		sw = append(sw, "m")
	}
	if e.Switches.Ungreedy {
		sw = append(sw, "U")
	}

	if len(sw) > 0 {
		e.Regex = "(?" + strings.Join(sw, "") + ")" + e.Regex
	}
}

func readExistingFlags(e *EditorArgs) {
	existingFlags := existingFlagsRx.FindStringSubmatch(e.Regex)
	if existingFlags == nil {
		return
	}
	e.Regex = existingFlagsRx.ReplaceAllString(e.Regex, "")

	onFlags := existingFlags[1]
	offFlags := existingFlags[3]

	if strings.Contains(onFlags, "i") {
		e.Switches.CaseSensitive = false
	}
	if strings.Contains(onFlags, "s") {
		e.Switches.SingleLine = true
	}
	if strings.Contains(onFlags, "m") {
		e.Switches.MultiLine = true
	}
	if strings.Contains(onFlags, "U") {
		e.Switches.Ungreedy = true
	}
	if strings.Contains(offFlags, "i") {
		e.Switches.CaseSensitive = true
	}
	if strings.Contains(offFlags, "s") {
		e.Switches.SingleLine = false
	}
	if strings.Contains(offFlags, "m") {
		e.Switches.MultiLine = false
	}
	if strings.Contains(offFlags, "U") {
		e.Switches.Ungreedy = false
	}
}

func (e *EditorArgs) GetArguments(argDefs []Arg, args []string) {
	e.OriginalArgs = args
	for ii, arg := range args {
		args[ii] = strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(
						strings.ReplaceAll(arg,
							"\\t", "\t"),
						"\\r", "\r"),
					"\\n", "\n"),
				"\\p", "|"),
			"\\q", "\"")
	}
	for ia, argDef := range argDefs {
		if len(args) > ia {
			argDef.SetValue(e, args[ia])
		} else {
			argDef.SetValue(e, argDef.DefaultValue)
		}
	}
}

func BuildArguments(cargs []Arg) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		for _, carg := range cargs {
			if carg.Position == len(args) {
				if len(carg.Options) > 0 {
					return carg.Options, cobra.ShellCompDirectiveNoFileComp
				}

				if carg.GetOptions != nil {
					opts := carg.GetOptions()
					if len(opts) > 0 {
						return opts, cobra.ShellCompDirectiveNoFileComp
					}
				}

				if carg.HelpText != "" {
					return cobra.AppendActiveHelp(
						nil,
						carg.HelpText), cobra.ShellCompDirectiveNoFileComp
				}

			}
		}

		return nil, cobra.ShellCompDirectiveNoFileComp
	}
}
