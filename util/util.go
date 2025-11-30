package util

import (
	"fmt"
)

func HelpE() error {
	fmt.Printf("commands: rep,grep,base64")
	return nil
}

func ListFunctions(fname string) {
	fmt.Printf("rep\ngrep\nbase64")
}
