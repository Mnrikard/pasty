package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Notify(input string) {
	fmt.Println(input)

	if !IsInputPiped() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nPress any key to continue")
		reader.ReadRune()
	}
}

func DisplayError(input error) {
	l := log.New(os.Stderr, "", 0)
	l.Println(input)
}
