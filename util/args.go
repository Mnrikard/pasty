package util

import "github.com/spf13/cobra"

type Arg struct {
	Position int
	HelpText string
	Options []string
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
