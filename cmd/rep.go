package cmd

import (
	"fmt"

	"github.com/mattr/pasty/text"
	"github.com/spf13/cobra"
)


var repCmd = &cobra.Command{
	Use:   "rep",
	Short: "replace text",
	Args: cobra.RangeArgs(1,2),
	Run: func(cmd *cobra.Command, args []string) {
		Replace(args[0], args[1]);
	},
}

func Replace(regex string, repla string) {
	text, err := text.GetText();
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(text);
}
