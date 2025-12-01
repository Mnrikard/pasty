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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(completionCmd)

	rootCmd.AddCommand(ColumnAlign.Command)
	rootCmd.AddCommand(Counter.Command)
	rootCmd.AddCommand(Dedup.Command)
	rootCmd.AddCommand(Lower.Command)
	rootCmd.AddCommand(Replace.Command)
	rootCmd.AddCommand(Title.Command)
	rootCmd.AddCommand(Upper.Command)
}
