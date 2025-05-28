package store

import "github.com/rizkikadafi/task-tracker/entity"

type TaskStore interface {
	LoadTasks() ([]entity.Task, error)
	SaveTasks([]entity.Task) error
}

