package edit

import (
	"encoding/base64"
	"html"
	"net/url"
)

func (e *EditorArgs) EncodeBase64(input string) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(input)), nil
}

func (e *EditorArgs) DecodeBase64(input string) (string, error) {
	output, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func (e *EditorArgs) EncodeForUrl(input string) (string, error) {
	return url.QueryEscape(input), nil
}

func (e *EditorArgs) DecodeFromUrl(input string) (string, error) {
	return url.QueryUnescape(input)
}

func (e *EditorArgs) EncodeForXml(input string) (string, error) {
	return html.EscapeString(input), nil
}

func (e *EditorArgs) DecodeFromXml(input string) (string, error) {
	return html.UnescapeString(input), nil
}

