package text

import (
	"fmt"
	"io"
	"os"

	"github.com/Mnrikard/pasty/edit"
	"github.com/Mnrikard/pasty/util"
	"github.com/tiagomelo/go-clipboard/clipboard"
)

var mockedText = ""
var mockedError error = nil
var isMocked = false

func SetMockedText(text string, err error) {
	isMocked = true
	mockedText = text
	mockedError = err
}

func GetMockedText() string {
	return mockedText
}

func GetText() (string, error) {
	if isMocked {
		return mockedText, mockedError
	}

	if util.IsInputPiped() {
		input, err := getStdInput()
		if err != nil {
			return "", err
		}
		return string(input), nil
	}

	c := clipboard.New()
	text, err := c.PasteText()
	if err != nil {
		return "", err
	}

	return text, nil
}

func SetText(input string) error {
	if isMocked {
		mockedText = input
		return mockedError
	}

	if util.IsInputPiped() {
		fmt.Print(input)
	} else {
		c := clipboard.New()
		if err := c.CopyText(input); err != nil {
			return err
		}
	}

	return nil
}

func getStdInput() ([]byte, error) {
	if !util.IsInputPiped() {
		return nil, nil
	}
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func EditText(e *edit.EditorArgs, fx func(string) (string, error)) {
	txt, err := GetText()
	if err != nil {
		util.DisplayError(err)
		return
	}

	newText, err := fx(txt)
	if err != nil {
		util.DisplayError(err)
		return
	}

	err = SetText(newText)
	if err != nil {
		util.DisplayError(err)
		return
	}
}
