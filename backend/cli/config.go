package cli

import "github.com/urfave/cli"

// ConfigureInfo configures global info about the cli
func ConfigureInfo(app *cli.App) {
	app.Name = "Simplest task manager"
	app.Usage = "./taskr command [command options] [arguments...]"
	app.Version = "0.0.1"
	app.Version = "Ilkin Musayev"
}
