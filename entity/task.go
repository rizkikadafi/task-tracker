package entity

import "time"

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTask(id int, title string) Task {
	now := time.Now()
	return Task{
		ID:        id,
		Title:     title,
		Status:    StatusTodo,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
