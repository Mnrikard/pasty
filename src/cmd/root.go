package cmd

import (
	"fmt"
	"os"

	"github.com/Mnrikard/pasty/edit"
	"github.com/Mnrikard/pasty/switches"
	"github.com/Mnrikard/pasty/text"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pasty",
	Short: "Pasty is clipboard editor to add macros to any application",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootSwitches = switches.Switches{}

func defineRegexSwitches(commandName string) {
	cmd := edit.FindCommandByName(rootCmd, commandName)
	if cmd == nil {
		return
	}

	cmd.Flags().BoolVarP(&rootSwitches.MultiLine, "multi-line", "m", false, "regex ^ and $ matches beginning and end of lines")
	cmd.Flags().BoolVarP(&rootSwitches.SingleLine, "single-line", "s", false, "regex . matches \\n")
	cmd.Flags().BoolVarP(&rootSwitches.CaseSensitive, "case-sensitive", "I", false, "regex is case sensitive")
	cmd.Flags().BoolVarP(&rootSwitches.Ungreedy, "un-greedy", "U", false, "regex patterns behave ungreedily")
}

func init() {
	cobra.MousetrapHelpText = ""
	rootCmd.AddCommand(completionCmd)

	for _, sc := range edit.SubCommands {
		if sc.Name == "udf" {
			sc.EditFunc = func(e *edit.EditorArgs) func(string) (string, error) { return e.ExecuteUdf }
		}
		if sc.Name == "plugin" {
			sc.EditFunc = func(e *edit.EditorArgs) func(string) (string, error) { return e.HandlePlugin }
		}
		cmd := buildCommand(sc)
		if sc.Name == "udf" {
			cmd.ValidArgs = edit.ListUdfs()
		}
		if sc.Name == "plugin" {
			cmd.ValidArgs = edit.ListPlugins()
		}
		rootCmd.AddCommand(cmd)
	}

	defineRegexSwitches("rep")
	defineRegexSwitches("grep")
	grep := edit.FindCommandByName(rootCmd, "grep")
	if grep != nil {
		grep.Flags().BoolVarP(&rootSwitches.GrepOnlyMatching, "only-matching", "o", false, "Returns only the matched (non-empty) parts of a matching line, with each such part on a separate output section")
		grep.Flags().BoolVarP(&rootSwitches.GrepInvertMatch, "invert-match", "v", false, "Invert the sense of matching, to select non-matching lines")
	}

	sorter := edit.FindCommandByName(rootCmd, "sort")
	if sorter != nil {
		sorter.Flags().BoolVarP(&rootSwitches.Invert, "invert", "v", false, "Sorts in descending order instead of ascending")
	}
}

func buildCommand(sc edit.SubCommand) *cobra.Command {
	cmd := &cobra.Command{
		Use:     sc.Use,
		Short:   sc.Short,
		Long:    sc.Long,
		Aliases: sc.Aliases,
		Args:    sc.Args,
	}

	if len(sc.ArgDefs) > 0 {
		cmd.ValidArgsFunction = edit.BuildArguments(sc.ArgDefs)
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		e.Switches = &rootSwitches
		if len(sc.ArgDefs) > 0 {
			e.GetArguments(sc.ArgDefs, args)
		}

		if sc.CustomSetup != nil {
			sc.CustomSetup(cmd, e)
		}

		text.EditText(e, sc.EditFunc(e))
	}

	return cmd
}
