package edit

import (
	"encoding/json"
	"net/url"
	"strings"
)

func (e *EditorArgs) ParseUrl(input string) (string, error) {
	url, err := url.Parse(strings.Trim(input, " 	\r\n"))
	if err != nil {
		return "", err
	}

	output := make(map[string]interface{})
	output["scheme"] = url.Scheme
	if url.User != nil {
		output["user"] = url.User
	}
	output["host"] = url.Host
	output["path"] = url.Path
	output["rawQuery"] = url.RawQuery
	output["query"] = url.Query()
	if url.RawFragment != "" {
		output["fragment"] = url.RawFragment
	}
/*
"ForceQuery": false,
        "Fragment": "",
        "Host": "google.com",
        "OmitHost": false,
        "Opaque": "",
        "Path": "/query/search",
        "RawFragment": "",
        "RawPath": "",
        "RawQuery": "a=b\u0026c=http%2d%2f%2fsomewhere",
        "Scheme": "https",
        "User": null,
        "query": {
                "a": [
                        "b"
                ],
                "c": [
                        "http-//somewhere"
                ]
        }
*/

	outb, err := json.MarshalIndent(output, "", settings().TabString)
	if err != nil {
		return "", err
	}

	return string(outb), nil
}
