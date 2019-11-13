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
	Name     string
	Deadline string
}

// TaskSection - a single section of tasks
type TaskSection struct {
	Section section
	Tasks   []Task
}

// TaskBoard a single board of tasks
type TaskBoard struct {
	Backlog  TaskSection
	Dev      TaskSection
	Progress TaskSection
	Done     TaskSection
}
