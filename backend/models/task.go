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

// Section of the task
type section string

// Task model
type Task struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Deadline string `json:"deadline"`
}

// TaskSection - a single section of tasks
type TaskSection struct {
	Tasks []Task `json:"tasks"`
}

// FindTask finds the task and returns the index if found or -1 if not found
func (section *TaskSection) FindTask(task string) int {
	index := -1
	for i, item := range section.Tasks {
		if item.ID == task {
			index = i
			break
		}
	}
	return index
}

// RemoveTask removes the task by the index
func (section *TaskSection) RemoveTask(index int) {
	section.Tasks = append(section.Tasks[:index], section.Tasks[index+1:]...)
}

// TaskBoard a single board of tasks
type TaskBoard struct {
	Backlog  TaskSection `json:"backlog"`
	Dev      TaskSection `json:"dev"`
	Progress TaskSection `json:"progress"`
	Done     TaskSection `json:"done"`
}
