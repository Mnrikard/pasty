package cmd

import (
	"github.com/mattr/pasty/edit"
	"github.com/mattr/pasty/text"
	"github.com/spf13/cobra"
)

var SolveMath = &cobra.Command{
	Use:   "math",
	Short: "Evaluates simple math equations",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		text.EditText(e, e.SolveMath)
	},
}
