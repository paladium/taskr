package cli

import (
	"taskr/storage"

	"github.com/urfave/cli"
)

// Init function initialises the cli app with all nessesary commands
func Init() *cli.App {
	storage.Init()
	app := cli.NewApp()
	ConfigureInfo(app)
	ConfigureCommands(app)
	return app
}
