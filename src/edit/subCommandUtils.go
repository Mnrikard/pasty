package edit

import (
	"strings"

	"github.com/spf13/cobra"
)

func FindCommandByName(rootCmd *cobra.Command, name string) *cobra.Command {
	for _, c := range rootCmd.Commands() {
		if strings.EqualFold(c.Name(), name) {
			return c
		}
	}

	return nil
}

func FindSubCommandsByNameOrAlias(name string) *SubCommand {
	for _, c := range SubCommands {
		if strings.EqualFold(name, c.Name) {
			return &c
		}
	}

	for _, aliasC := range SubCommands {
		for _, a := range aliasC.Aliases {
			if strings.EqualFold(name, a) {
				return &aliasC
			}
		}
	}

	return nil
}
