package text

import (
	"io"
	"log"
	"os"

	"github.com/golang-design/clipboard"
)


func GetText() (string, error) {
	if isInputPiped() {
		return string(getStdInput()), nil;
	}

	err := clipboard.init()
	if err != nil {
		return "", err
	}

	return clipboard.Read(clipboard.FmtText)
}

func getStdInput() []byte {
	if !isInputPiped() {
		return nil
	}
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func isInputPiped() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
