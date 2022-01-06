package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Shell(config map[string]string) {

	input := bufio.NewScanner(os.Stdin)
	fmt.Print(config["shellPrompt"])

	for input.Scan() {

		cmd := input.Text()
		if strings.Compare(cmd, "") != 0 {
			switch cmd {
			case "help":
				Help()
				break
			case "exit":
				//Todo : What happen with CTRL+C ?
				os.Exit(0)
			case "show":
				ShowConfig(config)
				break
			default:
				fmt.Println(" Command not found >" + cmd + "<")
				break
			}
		}
		fmt.Print(config["shellPrompt"])
	}
}
