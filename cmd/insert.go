package cmd

import (
	"github.com/mattr/pasty/edit"
	"github.com/mattr/pasty/text"
	"github.com/spf13/cobra"
)

var insertArgDefs = []edit.Arg{
	{
		Position: 0,
		HelpText: "Table Name",
		SetValue: func (e *edit.EditorArgs, value string) {
			e.Option = value
		},
	},
	{
		Position: 1,
		HelpText: "Column Delimiter",
		SetValue: func (e *edit.EditorArgs, value string) {
			e.ColumnDelimiter = value
		},
	},
	{
		Position: 2,
		HelpText: "Row Delimiter",
		SetValue: func (e *edit.EditorArgs, value string) {
			e.RowDelimiter = value
		},
		DefaultValue: "\r?\n",
	},
}

var insertSql = &cobra.Command{
	Use:   "insert",
	Short: "Converts result sets into an insert statement",
	Args:  cobra.MinimumNArgs(2),
	ValidArgsFunction: edit.BuildArguments(insertArgDefs),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		e.GetArguments(insertArgDefs, args)
		text.EditText(e, e.InsertSQL)
	},
}
