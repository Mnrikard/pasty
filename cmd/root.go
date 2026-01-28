package cmd

import (
	"fmt"
	"os"

	"github.com/mattr/pasty/switches"
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

func DefineRegexSwitches(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&rootSwitches.MultiLine,     "multi-line",     "m", false, "regex ^ and $ matches beginning and end of lines")
	cmd.Flags().BoolVarP(&rootSwitches.SingleLine,    "single-line",    "s", false, "regex . matches \\n")
	cmd.Flags().BoolVarP(&rootSwitches.CaseSensitive, "case-sensitive", "I", false, "regex is case sensitive")
	cmd.Flags().BoolVarP(&rootSwitches.Ungreedy,      "un-greedy",      "U", false, "regex patterns behave ungreedily")
}

func init() {

	rootCmd.AddCommand(completionCmd)

	rootCmd.AddCommand(ColumnAlign)
	rootCmd.AddCommand(Counter)
	rootCmd.AddCommand(Dedup)
	rootCmd.AddCommand(Lower)
	rootCmd.AddCommand(Replace)
	DefineRegexSwitches(Replace)
	rootCmd.AddCommand(Title)
	rootCmd.AddCommand(Upper)
	rootCmd.AddCommand(Base64Encode)
	rootCmd.AddCommand(Base64Decode)
	rootCmd.AddCommand(UrlEncode)
	rootCmd.AddCommand(UrlDecode)
	rootCmd.AddCommand(XmlEncode)
	rootCmd.AddCommand(XmlDecode)
	rootCmd.AddCommand(Grep)
	rootCmd.AddCommand(insertSql)
	rootCmd.AddCommand(SetText)
	rootCmd.AddCommand(NewGuid)
	rootCmd.AddCommand(FromBase)
	rootCmd.AddCommand(ToBase)
	rootCmd.AddCommand(SolveMath)
	DefineRegexSwitches(Grep)
	Grep.Flags().BoolVarP(&rootSwitches.GrepOnlyMatching, "only-matching", "o", false, "Returns only the matched (non-empty) parts of a matching line, with each such part on a separate output section")
	Grep.Flags().BoolVarP(&rootSwitches.GrepInvertMatch, "invert-match", "v", false, "Invert the sense of matching, to select non-matching lines")
}
