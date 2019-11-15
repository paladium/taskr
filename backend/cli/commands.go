package cli

import (
	"fmt"
	"log"
	"net/http"
	"taskr/config"
	"taskr/models"
	"taskr/output"
	"taskr/server"
	"taskr/storage"

	"github.com/urfave/cli"
)

// ConfigureCommands defines the commands for the cli
func ConfigureCommands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Inits the task board in the current directory",
			Action: func(c *cli.Context) {
				storage.GetStorage().SaveBoard(&models.TaskBoard{})
				fmt.Println("Project succesfully created ✅")
			},
		},
		{
			Name:    "backlog",
			Aliases: []string{"b"},
			Usage:   "Add new task to a backlog",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "deadline, d", Required: false},
				cli.StringFlag{Name: "task,t", Required: true},
			},
			Action: func(c *cli.Context) {
				taskName := c.String("task")
				deadline := c.String("deadline")
				storage.GetStorage().AddTask(&models.Task{
					Name:     taskName,
					Deadline: deadline,
					ID:       storage.GetRandomName(0),
				})
				fmt.Println("Task added succesfully 🚀")
			},
		},
		{
			Name:    "tasks",
			Aliases: []string{"t"},
			Usage:   "List all the tasks",
			Action: func(c *cli.Context) {
				board := storage.GetStorage().GetBoard()
				output.Board(board)
			},
		},
		{
			Name:    "move",
			Aliases: []string{"mv"},
			Usage:   "Move the task along the board",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "task,t", Required: true},
			},
			Action: func(c *cli.Context) {
				task := c.String("task")
				storage.GetStorage().MoveTask(task)
				board := storage.GetStorage().GetBoard()
				output.Board(board)
			},
		},
		{
			Name:    "ui",
			Aliases: []string{"u"},
			Usage:   "Launches the server and ui for visual task managment",
			Action: func(c *cli.Context) {
				config.Init("dev")
				//Start the ui as the server
				go func() {
					box := config.GetBox()
					http.Handle("/", http.FileServer(box))
					log.Printf("Started the ui server at %s\n", config.GetConfig().GetString("server.ui"))
					http.ListenAndServe(config.GetConfig().GetString("server.ui"), nil)
				}()
				server.Init()
			},
		},
	}
}
