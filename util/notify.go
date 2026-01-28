package util

import (
	"bufio"
	"fmt"
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
	fmt.Errorf("%w", input)
}
