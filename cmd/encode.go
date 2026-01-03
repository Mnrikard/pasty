package cmd

import (
	"github.com/mattr/pasty/edit"
	"github.com/mattr/pasty/text"
	"github.com/spf13/cobra"
)

var Base64Encode = &cobra.Command{
	Use:   "base64encode",
	Aliases: []string{"base64"},
	Short: "Base64 encodes the text",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		text.EditText(e, e.EncodeBase64)
	},
}

var Base64Decode = &cobra.Command{
	Use:   "base64decode",
	Short: "Base64 decodes the text",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		text.EditText(e, e.DecodeBase64)
	},
}

var UrlEncode = &cobra.Command{
	Use:   "urlencode",
	Aliases: []string{"url"},
	Short: "Url encodes the text",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		text.EditText(e, e.EncodeForUrl)
	},
}

var UrlDecode = &cobra.Command{
	Use:   "urldecode",
	Short: "Url decodes the text",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		text.EditText(e, e.DecodeFromUrl)
	},
}

var XmlEncode = &cobra.Command{
	Use:   "xmlencode",
	Short: "XML encodes the text",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		text.EditText(e, e.EncodeForXml)
	},
}

var XmlDecode = &cobra.Command{
	Use:   "xmldecode",
	Short: "XML decodes the text",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		e := &edit.EditorArgs{}
		text.EditText(e, e.DecodeFromXml)
	},
}
