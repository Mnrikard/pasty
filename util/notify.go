package util

import (
	_ "embed"

	"github.com/gen2brain/beeep"
)

//go:embed binaries/pasty.png
var pastyIcon []byte

func Notify(title, message string) {
	beeep.AppName = "pasty"
	err := beeep.Notify(title, message, pastyIcon)
	if err != nil {
		panic(err)
	}
}

func DisplayError(message error) {
	Notify("Error", message.Error())
}
