package forms

// MoveTask struct for parsing json
type MoveTask struct {
	ID string `json:"id" binding:"required"`
}

// AddTask struct for parsing json
type AddTask struct {
	Name     string `json:"name" binding:"required"`
	Deadline string `json:"deadline"`
}
