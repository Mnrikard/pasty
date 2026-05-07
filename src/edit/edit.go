package edit

import (
	"strings"

	"github.com/Mnrikard/pasty/switches"
	"github.com/spf13/cobra"
)

type Arg struct {
	Position     int
	HelpText     string
	Options      []string
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
	Invert          bool
	OriginalArgs    []string
	Switches        *switches.Switches
}

func (e *EditorArgs) PrependRegex(rootSwitches switches.Switches) {
	sw := make([]string, 0)
	if !rootSwitches.CaseSensitive {
		sw = append(sw, "i")
	}
	if rootSwitches.SingleLine {
		sw = append(sw, "s")
	}
	if rootSwitches.MultiLine {
		sw = append(sw, "m")
	}
	if rootSwitches.Ungreedy {
		sw = append(sw, "U")
	}

	if len(sw) > 0 {
		e.Regex = "(?" + strings.Join(sw, "") + ")" + e.Regex
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
				if carg.HelpText != "" {
					return cobra.AppendActiveHelp(
						nil,
						carg.HelpText), cobra.ShellCompDirectiveNoFileComp
				}

				if len(carg.Options) > 0 {
					return carg.Options, cobra.ShellCompDirectiveNoFileComp
				}
			}
		}

		return nil, cobra.ShellCompDirectiveNoFileComp
	}
}
