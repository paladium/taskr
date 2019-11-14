package storage

import (
	"encoding/json"
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

// SaveTask saves a new task to the needed section
func (storage *TaskStorage) SaveTask(task *models.Task, section string) {
	board := storage.GetBoard()
	if section == models.BACKLOG {
		board.Backlog.Tasks = append(board.Backlog.Tasks, *task)
	}
	storage.SaveBoard(board)
}

// SaveBoard saves the new board to the file
func (storage *TaskStorage) SaveBoard(board *models.TaskBoard) {
	file, _ := json.MarshalIndent(board, "", " ")
	ioutil.WriteFile(getFileLocation(), file, 0666)
}

func (storage *TaskStorage) MoveTask(task string, section string) {
	board := storage.GetBoard()
	// currentSection
}

// GetStorage returns the current task storage
func GetStorage() *TaskStorage {
	return storage
}
