package edit

import (
	"os"
	"strings"

	"github.com/Mnrikard/pasty/util"
	"github.com/spf13/cobra"
)

var notify = util.Notify

var settings = util.Settings

var osUserHomeDir = os.UserHomeDir
var osStat = os.Stat
var osIsNotExist = os.IsNotExist
var osCreate = os.Create
var osReadFile = os.ReadFile
var osReadDir = os.ReadDir

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
