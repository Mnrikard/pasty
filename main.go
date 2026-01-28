package main

import (
	"github.com/mattr/pasty/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.EnableCaseInsensitive = true
	cmd.Execute()
}
