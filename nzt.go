package main

import (
	"flag"
	nztCommands "github.com/darknautic/nzt/commands"
	"log"
	"os"
	"strings"
)

var nztHome, isNztHomeSet = os.LookupEnv("NZT_HOME")
var nztTagDelimiter, isNztTagDelimiterSet = os.LookupEnv("NZT_TAG_DELIMITER")

var config = map[string]string{
	"shellPrompt":     "nzt > ",
	"nztHome":         nztHome,
	"nztTagDelimiter": nztTagDelimiter,
}

func main() {

	if isNztHomeSet {
		if len(os.Args) >= 2 {

			addFlag := flag.String("a", "", "Add a note")
			getFlag := flag.String("g", "", "Get a note")
			tagFlag := flag.String("t", "", "Tag(s)")
			shellFlag := flag.Bool("s", false, "Interactive shell")

			flag.Parse()

			addNote := *addFlag
			getNote := *getFlag
			tags := *tagFlag
			isShellFlag := *shellFlag

			if isShellFlag {
				nztCommands.Shell(config)

			} else {
				tagList := []string{"NA"}
				if len(tags) > 0 {
					if nztTagDelimiter == " " {
						tagIndex := 0
						for i, n := range os.Args {
							if n == "-t" {
								tagIndex = i
								break
							}
						}
						tagList = os.Args[(tagIndex + 1):]

					} else {
						tagList = strings.Split(tags, nztTagDelimiter)
					}
				}

				if len(addNote) > 0 {
					nztCommands.AddNote(addNote, tagList)
				} else {
					if len(getNote) > 0 {
						nztCommands.GetNote(getNote, tagList)
					} else {
						log.Fatal("No Add/Get option given.")
					}
				}

			}

		} else {
			nztCommands.Help()
		}

	} else {
		log.Fatal("Environment variable NZT_HOME is not set")
		os.Exit(1)
	}

}
