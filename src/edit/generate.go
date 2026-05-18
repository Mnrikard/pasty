package edit

import "github.com/google/uuid"

func (e *EditorArgs) SetText(input string) (string, error) {
	return e.Option, nil
}

func (e *EditorArgs) NewGuid(input string) (string, error) {
	return uuid.New().String(), nil
}
