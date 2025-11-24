package text

import (
	"fmt"

	"github.com/tiagomelo/go-clipboard/clipboard"
)

func SetText(input string) error {
	if isInputPiped() {
		fmt.Println(input);
	} else {
		c := clipboard.New()
		if err := c.CopyText(input); err != nil {
			return err
		}
	}

	return nil
}
