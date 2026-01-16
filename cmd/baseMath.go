package cmd

import (
	"github.com/mattr/pasty/edit"
	"github.com/mattr/pasty/text"
	"github.com/spf13/cobra"
)

var baseMathArgDefs = []edit.Arg{
	{
		Position: 0,
		HelpText: "Base",
		SetValue: func (e *edit.EditorArgs, value string) {
			e.Option = value
		},
	},
}

var ToBase = &cobra.Command{
	Use:   "toBase",
	Short: "converts base 10 to the given base",
	Args:  cobra.MinimumNArgs(1),
	ValidArgsFunction: edit.BuildArguments(baseMathArgDefs),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		e.GetArguments(baseMathArgDefs, args)
		text.EditText(e, e.ToNumBase)
	},
}

var FromBase = &cobra.Command{
	Use:   "fromBase",
	Short: "converts from the given base to base 10",
	Args:  cobra.MinimumNArgs(1),
	ValidArgsFunction: edit.BuildArguments(baseMathArgDefs),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		e.GetArguments(baseMathArgDefs, args)
		text.EditText(e, e.FromNumBase)
	},
}
