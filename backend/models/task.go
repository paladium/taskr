package models

const (
	// BACKLOG group
	BACKLOG = "backlog"
	//DEV group
	DEV = "dev"
	// PROGRESS group
	PROGRESS = "progress"
	//DONE group
	DONE = "done"
)

// Task model
type Task struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Deadline string `json:"deadline"`
	Section  string `json:"section" binding:"required"`
}

// TaskSection a slice of tasks
type TaskSection []Task

// TaskBoard - a single section of tasks
type TaskBoard struct {
	Tasks []Task `json:"tasks"`
}

// FindTask finds the task and returns the index if found or -1 if not found
func (board *TaskBoard) FindTask(task string) int {
	index := -1
	for i, item := range board.Tasks {
		if item.ID == task {
			index = i
			break
		}
	}
	return index
}

// RemoveTask removes the task by the index
func (board *TaskBoard) RemoveTask(index int) {
	board.Tasks = append(board.Tasks[:index], board.Tasks[index+1:]...)
}

// TaskBoard a single board of tasks

func inSection(section string) func(task Task) bool {
	return func(task Task) bool {
		return task.Section == section
	}
}

func filter(tasks []Task, test func(Task) bool) (filtered []Task) {
	for _, item := range tasks {
		if test(item) {
			filtered = append(filtered, item)
		}
	}
	return
}

// Backlog function returns the slice of the tasks only in backlog
func (board *TaskBoard) Backlog() TaskSection {
	return filter(board.Tasks, inSection(BACKLOG))
}

// Dev function returns the slice of the tasks only in development
func (board *TaskBoard) Dev() TaskSection {
	return filter(board.Tasks, inSection(DEV))
}

// Progress function returns the slice of the tasks only in development
func (board *TaskBoard) Progress() TaskSection {
	return filter(board.Tasks, inSection(PROGRESS))
}

// Done function returns the slice of the tasks only in development
func (board *TaskBoard) Done() TaskSection {
	return filter(board.Tasks, inSection(DONE))
}
