package commands

import (
	b64 "encoding/base64"
	"fmt"
)

func AddNote(text string, tags []string) {
	clearText := text
	encodedText := b64.StdEncoding.EncodeToString([]byte(clearText))
	fmt.Println("Adding a note ... >", clearText, "<")
	fmt.Println("ENCODED ... >", encodedText, "<")
	decodedText, _ := b64.StdEncoding.DecodeString(encodedText)
	fmt.Println("DECODED ... >", string(decodedText), "<")

	for _, tag := range tags {
		fmt.Println("tag >", tag, "<")
	}
}
