package taskmanager

import "time"

// Priority enum (Extendable)
type Priority int

const (
	Low Priority = iota
	Medium
	High
)

func (p Priority) String() string {
	return [...]string{"LOW", "MEDIUM", "HIGH"}[p]
}

// Status enum
type Status int

const (
	Open Status = iota
	InProgress
	Completed
)

func (s Status) String() string {
	return [...]string{"OPEN", "IN_PROGRESS", "COMPLETED"}[s]
}

// Task entity
// (Single Responsibility: only holds data)
type Task struct {
	ID          string
	Title       string
	Description string
	Priority    Priority
	DueDate     time.Time
	CreatedDate time.Time
	Status      Status
}

// NewTask is a Factory Method for Task creation
// (Factory Pattern)
func NewTask(id, title, desc string, priority Priority, due time.Time) *Task {
	return &Task{
		ID:          id,
		Title:       title,
		Description: desc,
		Priority:    priority,
		DueDate:     due,
		CreatedDate: time.Now(),
		Status:      Open,
	}
}
