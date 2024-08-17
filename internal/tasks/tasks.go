package tasks

import "fmt"

// Task represents a single todo-item
type Task struct {
	ID          int
	Description string
}

// New creates an intance of a task and returns the pointer to that task
func New(id int, description string) *Task {
	return &Task{ID: id, Description: description}
}

// String returns a string representation of the task
func (t *Task) String() string {
	return fmt.Sprintf("ID: %d\t%s", t.ID, t.Description)
}
