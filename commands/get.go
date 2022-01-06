package commands

import "fmt"

func GetNote(text string, tags []string) {
	fmt.Println("Getting a note ... >", text, "<")
	for _, tag := range tags {
		fmt.Println("tag >", tag, "<")
	}
}
