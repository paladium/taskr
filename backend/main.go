package main

import (
	"log"
	"os"
	"taskr/cli"
)

func main() {
	// home, err := os.UserHomeDir()
	// if err != nil {
	// 	fmt.Errorf("Cannot find the home directory")
	// }
	app := cli.Init()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
