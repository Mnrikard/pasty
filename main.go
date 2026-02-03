package main

import (
 "os"
	"os/exec"
	"runtime"
	"golang.org/x/term"
	"github.com/mattr/pasty/cmd"
	"github.com/spf13/cobra"
)

func main() {
	if runtime.GOOS == "windows" && !term.IsTerminal(int(os.Stdout.Fd())) {
		cmd := exec.Command("cmd", "/k", os.Args[0])
		cmd.Run()
		return 
	}

 cobra.EnableCaseInsensitive = true
	cmd.Execute()
}
