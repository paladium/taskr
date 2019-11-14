package main

import (
	"log"
	"os"
	"taskr/cli"
)

func main() {
	app := cli.Init()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
