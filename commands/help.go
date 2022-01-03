package commands

import "fmt"

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
 * 		- a, add <text>
 *			add a note
 *
 *		- t, tags <tag-1 ...tag-n>
 *			space-separated list of tags
 *
 *		- l , list <config|tags|datafiles>
 *
 *		- s , shell
 *			get into interactive shell mode
 *
 *		- f , find
 *			shows notes whose labels match filter
 *
 *		- h , help
 *			display this message
 *
 *		- exit
 *
 * DESCRIPTION
 *
 * 	features
 *		- zettelkasten
 *
 *
	`)
}
