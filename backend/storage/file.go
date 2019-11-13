package storage

import (
	"taskr/models"
)

// TaskStorage for saving and retrieving tasks
type TaskStorage struct{}

var storage *TaskStorage

// Init file-based storage
func Init() {
	storage = new(TaskStorage)
}

// SaveTask saves a new task to the needed section
func (storage *TaskStorage) SaveTask(task *models.Task) {

}

// GetStorage returns the current task storage
func GetStorage() *TaskStorage {
	return storage
}
