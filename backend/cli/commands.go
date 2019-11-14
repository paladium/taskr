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
			},
			Action: func(c *cli.Context) {
				taskName := c.Args().First()
				deadline := c.String("deadline")
				storage.GetStorage().SaveTask(&models.Task{
					Name:     taskName,
					Deadline: deadline,
				}, models.BACKLOG)
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
					testBoard := &models.TaskBoard{
						Backlog: models.TaskSection{Tasks: []models.Task{
							models.Task{Name: "Work on this", Deadline: "tomorrow"},
						}},
						Dev: models.TaskSection{Tasks: []models.Task{
							models.Task{Name: "Work on this", Deadline: "tomorrow"},
							models.Task{Name: "Work on this", Deadline: "tomorrow"},
							models.Task{Name: "Work on this", Deadline: "tomorrow"},
						}},
						Progress: models.TaskSection{Tasks: []models.Task{
							models.Task{Name: "Work on this", Deadline: "tomorrow"},
							models.Task{Name: "Work on this", Deadline: "tomorrow"},
							models.Task{Name: "Work on this", Deadline: "tomorrow"},
							models.Task{Name: "Work on this", Deadline: "tomorrow"},
							models.Task{Name: "Work on this", Deadline: "tomorrow"},
							models.Task{Name: "Work on this", Deadline: "tomorrow"},
						}},
						Done: models.TaskSection{Tasks: []models.Task{
							models.Task{Name: "Work on this", Deadline: "tomorrow"},
							models.Task{Name: "Work on this", Deadline: "tomorrow"},
							models.Task{Name: "Work on this", Deadline: "tomorrow"},
						}},
					}
					output.Board(testBoard)
				}
			},
		},
	}
}
