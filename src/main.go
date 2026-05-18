package main

import (
	"github.com/Mnrikard/pasty/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.EnableCaseInsensitive = true
	cmd.Execute()
}
