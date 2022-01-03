package commands

import "fmt"

func ShowConfig(config map[string]string) {
	fmt.Println("----------------")
	fmt.Println("NZT_HOME : " + config["NZT_HOME"])
	fmt.Println("----------------")
}
