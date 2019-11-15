package controllers

import (
	"net/http"
	"taskr/forms"
	"taskr/models"
	"taskr/storage"

	"github.com/gin-gonic/gin"
)

// TaskController for common tasks actions
type TaskController struct{}

// Tasks return all the tasks
func (t *TaskController) Tasks(c *gin.Context) {
	tasks := storage.GetStorage().GetBoard()
	if tasks.Tasks == nil {
		tasks.Tasks = []models.Task{}
	}
	c.JSON(http.StatusOK, tasks)
}

// MoveTask moves the task further in the pipeline
func (t *TaskController) MoveTask(c *gin.Context) {
	taskForm := new(forms.MoveTask)
	c.Bind(taskForm)
	storage.GetStorage().MoveTask(taskForm.ID)
	c.JSON(http.StatusOK, struct{ message string }{message: "ok"})
}

// AddTask adds new task to backlog
func (t *TaskController) AddTask(c *gin.Context) {
	taskForm := new(forms.AddTask)
	c.Bind(taskForm)
	storage.GetStorage().AddTask(&models.Task{
		ID:       storage.GetRandomName(0),
		Name:     taskForm.Name,
		Section:  models.BACKLOG,
		Deadline: taskForm.Deadline,
	})
}
