package cmd

import (
	"github.com/mattr/pasty/edit"
	"github.com/mattr/pasty/text"
	"github.com/spf13/cobra"
)

var Upper = &cobra.Command{
	Use:   "upper",
	Short: "Upper cases text",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		text.EditText(e, e.Upper)
	},
}

var Lower = &cobra.Command{
	Use:   "lower",
	Short: "Lower cases text",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		text.EditText(e, e.Lower)
	},
}

var Title = &cobra.Command{
	Use:   "title",
	Short: "Title cases text",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		text.EditText(e, e.Title)
	},
}
