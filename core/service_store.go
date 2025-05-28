package core

import (
	"errors"
	"time"

	"github.com/rizkikadafi/task-tracker/entity"
	"github.com/rizkikadafi/task-tracker/store"
)

type StoreBackedService struct {
	store  store.TaskStore
	tasks  []entity.Task
	lastID int
}

// Factory function
func NewStoreService(s store.TaskStore) (*StoreBackedService, error) {
	tasks, err := s.LoadTasks()
	if err != nil {
		return nil, err
	}

	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	return &StoreBackedService{
		store:  s,
		tasks:  tasks,
		lastID: maxID,
	}, nil
}

func (s *StoreBackedService) AddTask(title string) (entity.Task, error) {
	s.lastID++
	task := entity.NewTask(s.lastID, title)
	s.tasks = append(s.tasks, task)

	if err := s.store.SaveTasks(s.tasks); err != nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (s *StoreBackedService) ListTasks() []entity.Task {
	return s.tasks
}

func (s *StoreBackedService) UpdateStatus(id int, status entity.Status) error {
	found := false
	for i, t := range s.tasks {
		if t.ID == id {
			s.tasks[i].Status = status
			s.tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		return errors.New("task not found")
	}

	return s.store.SaveTasks(s.tasks)
}

func (s *StoreBackedService) DeleteTask(id int) error {
	idx := -1
	for i, t := range s.tasks {
		if t.ID == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		return errors.New("task not found")
	}

	// hapus dari slice
	s.tasks = append(s.tasks[:idx], s.tasks[idx+1:]...)
	return s.store.SaveTasks(s.tasks)
}

