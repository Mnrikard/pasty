package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var pasteCmd = &cobra.Command{
	Use:   "paste",
	Short: "paste from clipboard",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the clipboard
		err := clipboard.Init()
		if err != nil {
			fmt.Printf("Error initializing clipboard: %v\n", err)
			return
		}

		content := clipboard.Read(clipboard.FmtText)
		if content == nil {
			fmt.Println("clipboard is empty")
			return
		}
		fmt.Println(string(content))
	},
}
