package util

import "os"

func IsInputPiped() bool {
	stat, err := os.Stdin.Stat()
	return err == nil && (stat.Mode() & os.ModeCharDevice) == 0
}
