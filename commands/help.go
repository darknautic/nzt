package commands

import (
	"fmt"
	"os"
)

func Help() {
	fmt.Println(`
 *
 * 
 * NAME
 *	nzt - cli note-taking app
 *
 * SYNOPSIS
 *
 *	nzt <command> <arg>
 *
 * 	commands
 * 		-a <text>
 *			add a note
 *		-t <tag-1 ...tag-n>
 *			list of tags
 *			Customize delimiter character with
 *			i.e NZT_TAG_DELIMITER=","
 *		-l <config|tags|datafiles>
 *			list
 *		-s
 *			get into interactive shell mode
 *			subcommands :
 *				- help
 *				- exit
 *		-h
 *			display command help
 *
 *
 * DESCRIPTION
 *
 * 	features
 *		- zettelkasten
 *
 *
	`)
	os.Exit(0)
}
