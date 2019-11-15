package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"taskr/models"
	"time"
)

// TaskStorage for saving and retrieving tasks
type TaskStorage struct{}

var storage *TaskStorage

// Init file-based storage
func Init() {
	storage = new(TaskStorage)
	//Setup random generator
	rand.Seed(time.Now().UnixNano())
}

func getFileLocation() string {
	home, _ := os.UserHomeDir()
	return home + "/" + ".taskr.json"
}

// GetBoard reads the file and returns the board
func (storage *TaskStorage) GetBoard() *models.TaskBoard {
	file, err := os.OpenFile(getFileLocation(), os.O_CREATE|os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		log.Fatal("Cannot create or open the file")
	}
	byteValue, _ := ioutil.ReadAll(file)
	var board models.TaskBoard
	json.Unmarshal(byteValue, &board)
	return &board
}

// AddTask saves a new task to the needed section
func (storage *TaskStorage) AddTask(task *models.Task) {
	board := storage.GetBoard()
	task.Section = models.BACKLOG
	board.Tasks = append(board.Tasks, *task)
	storage.SaveBoard(board)
}

// SaveBoard saves the new board to the file
func (storage *TaskStorage) SaveBoard(board *models.TaskBoard) {
	file, _ := json.MarshalIndent(board, "", " ")
	ioutil.WriteFile(getFileLocation(), file, 0666)
}

// MoveTask moves task further in the section
func (storage *TaskStorage) MoveTask(task string) {
	board := storage.GetBoard()
	taskIndex := board.FindTask(task)
	if taskIndex == -1 {
		fmt.Printf("Task '%s' cannot be found 😢\n", task)
		return
	}
	currentTask := &board.Tasks[board.FindTask(task)]
	switch currentTask.Section {
	case models.BACKLOG:
		currentTask.Section = models.DEV
	case models.DEV:
		currentTask.Section = models.PROGRESS
	case models.PROGRESS:
		currentTask.Section = models.DONE
	}
	storage.SaveBoard(board)
}

// GetStorage returns the current task storage
func GetStorage() *TaskStorage {
	return storage
}
