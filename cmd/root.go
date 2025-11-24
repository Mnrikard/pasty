package cmd

import (
	"fmt"
	"os"

	"github.com/mattr/pasty/util"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pasty",
	Short: "Pasty is clipboard editor to add macros to any application",
	RunE: func(cmd *cobra.Command, args []string) error {
		return util.HelpE()
	},
}



var caseCmd = &cobra.Command{
	Use:   "case",
	Short: "change case",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("case command")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(repCmd)
	rootCmd.AddCommand(caseCmd)
	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(pasteCmd)
}
