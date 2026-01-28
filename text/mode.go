package text

import "os"

func isInputPiped() bool {
	stat, err := os.Stdin.Stat()
	return err == nil && (stat.Mode() & os.ModeCharDevice) == 0
}
