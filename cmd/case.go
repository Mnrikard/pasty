package cmd

import (
	"strings"

	"github.com/mattr/pasty/text"
	"github.com/mattr/pasty/util"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var Upper = &util.Editor {
	Command: &cobra.Command{
		Use:   "upper",
		Short: "Upper cases text",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			changeCase(strings.ToUpper)
		},
	},
}

var Lower = &util.Editor {
	Command: &cobra.Command{
		Use:   "lower",
		Short: "Lower cases text",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			changeCase(strings.ToLower)
		},
	},
}

var Title = &util.Editor {
	Command: &cobra.Command{
		Use:   "title",
		Short: "Title cases text",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			changeCase(titleCase)
		},
	},
}

func changeCase(fx func (string)(string)) {
	txt, err := text.GetText();
	if err != nil {
		panic(err)
	}

	replacedText := fx(txt)

	err = text.SetText(replacedText)
	if err != nil {
		panic(err)
	}
}

func titleCase(input string) string {
	caser := cases.Title(language.English)
	return caser.String(input)
}
