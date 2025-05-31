package entity

import (
	"errors"
	"strings"
	"time"
)

type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in-progress"
	Done       Status = "done"
)

func ParseStatus(s string) (Status, error) {
	switch strings.ToLower(s) {
	case "todo":
		return Todo, nil
	case "in-progress":
		return InProgress, nil
	case "done":
		return Done, nil
	default:
		return "", errors.New("invalid status: must be todo, in-progress, or done")
	}
}

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
		Status:    Todo,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
