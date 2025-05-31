package core

import (
	"errors"
	"time"

	"github.com/rizkikadafi/task-tracker/entity"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type TaskService interface {
	AddTask(title string) (entity.Task, error)
	ListTasks() ([]entity.Task, error)
	FindTaskByID(id int) (entity.Task, error)
	UpdateStatus(id int, status entity.Status) error
	DeleteTask(id int) error
}

func NewTask(id int, title string) entity.Task {
	return entity.Task{
		ID:        id,
		Title:     title,
		Status:    entity.Todo,
		CreatedAt: time.Now(),
	}
}

type DefaultTaskService struct {
	tasks  []entity.Task
	lastID int
}

func NewTaskService() *DefaultTaskService {
	return &DefaultTaskService{
		tasks:  []entity.Task{},
		lastID: 0,
	}
}

func (s *DefaultTaskService) AddTask(title string) (entity.Task, error) {
	s.lastID++
	task := NewTask(s.lastID, title)
	s.tasks = append(s.tasks, task)
	return task, nil
}
