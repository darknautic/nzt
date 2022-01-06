package commands

import "fmt"

func ShowConfig(config map[string]string) {
	fmt.Println("----------------")
	fmt.Println("NZT_HOME : " + config["nztHome"])
	fmt.Println("NZT_TAG_DELIMITER : " + config["nztTagDelimiter"])
	fmt.Println("----------------")
}
