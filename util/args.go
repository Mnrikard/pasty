package util

import "github.com/spf13/cobra"

type Arg struct {
	Position int
	HelpText string
	Options []string
	SetValue func(*Editor, string)
	DefaultValue string
}

type Editor struct {
	Regex string
	Replacement string
	ColumnDelimiter string
	RowDelimiter string
	NumSpaces int
	Option string

	ArgDefs []Arg
	Command *cobra.Command
}

func GetArguments(e *Editor, argDefs []Arg, args []string) {
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
