package text

import (
	"io"
	"log"
	"os"

	"github.com/tiagomelo/go-clipboard/clipboard"
)

func GetText() (string, error) {
	if isInputPiped() {
		return string(getStdInput()), nil
	}

	log.Println("getting data from clipboard")

	c := clipboard.New()
	text, err := c.PasteText()
	if err != nil {
		return "", err
	}

	log.Printf("text: %q\n", text)

	return text, nil
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

