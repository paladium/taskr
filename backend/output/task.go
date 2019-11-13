package output

import (
	"os"
	"taskr/models"

	"github.com/jedib0t/go-pretty/table"
)

func getTaskOrEmpty(section models.TaskSection, index int) *models.Task {
	if len(section.Tasks) > index {
		return &section.Tasks[index]
	}
	return nil
}

func getTaskSlice(sections []models.TaskSection, index int) table.Row {
	slice := table.Row{}
	for _, item := range sections {
		task := getTaskOrEmpty(item, index)
		if task == nil {
			slice = append(slice, "")
		} else {
			slice = append(slice, task.Name)
		}
	}
	return slice
}

func appendTasks(board *models.TaskBoard, t table.Writer) func(models.TaskSection) {
	var i = 0
	return func(section models.TaskSection) {
		for i < len(section.Tasks) {
			slice := getTaskSlice([]models.TaskSection{board.Backlog, board.Dev, board.Progress, board.Done}, i)
			t.AppendRow(slice)
			i++
		}
	}
}

// Board prints a beatiful board in the console
func Board(board *models.TaskBoard) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredMagentaWhiteOnBlack)
	t.AppendHeader(table.Row{board.Backlog.Section, board.Dev.Section, board.Progress.Section, board.Done.Section})
	tableGenerator := appendTasks(board, t)
	tableGenerator(board.Backlog)
	tableGenerator(board.Dev)
	tableGenerator(board.Progress)
	tableGenerator(board.Done)
	t.Render()
}
