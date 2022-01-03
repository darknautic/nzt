package main

import (
	"flag"
	"fmt"
	nztCommands "github.com/darknautic/nzt/commands"
	"log"
	"os"
)

var NZT_HOME, ok = os.LookupEnv("NZT_HOME")

var config = map[string]string{
	"shellPrompt": "nzt > ",
	"NZT_HOME":    NZT_HOME,
}

func main() {

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	tagCmd := flag.NewFlagSet("tag", flag.ExitOnError)
	shellFlag := flag.Bool("shell", false, "interactive shell")

	if len(os.Args) == 2 {
		flag.Parse()
		isShellEnabled := *shellFlag
		if isShellEnabled {
			if !ok {
				//TODO , assess to use a default directory in home/user/.nzt
				log.Fatal("Environment variable NZT_HOME is not set")
				os.Exit(0)
			} else {
				nztCommands.ShowConfig(config)
				nztCommands.Shell(config)
			}
		}
		os.Exit(0)
	}

	if len(os.Args) < 4 {
		fmt.Println("Expected required flags 'add' and 'tag' ")
		os.Exit(1)
	}

	addCmd.Parse(os.Args[2:3])
	tagCmd.Parse(os.Args[4:])

	fmt.Println("adding ->", addCmd.Args())
	fmt.Println("tags -> ", tagCmd.Args())

}
