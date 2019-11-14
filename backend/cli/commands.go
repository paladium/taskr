package cli

import (
	"fmt"
	"taskr/models"
	"taskr/output"
	"taskr/storage"

	"github.com/urfave/cli"
)

// ConfigureCommands defines the commands for the cli
func ConfigureCommands(app *cli.App) {
	app.Commands = []cli.Command{
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
				storage.GetStorage().SaveTask(&models.Task{
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
			Flags: []cli.Flag{
				cli.StringFlag{Name: "section, s", Required: false},
			},
			Action: func(c *cli.Context) {
				section := c.String("section")
				if section == "" {
					board := storage.GetStorage().GetBoard()
					output.Board(board)
				}
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
	}
}
