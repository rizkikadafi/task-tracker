package jsonstore

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/rizkikadafi/task-tracker/entity"
)

type JSONStore struct {
	filePath string
	mu       sync.Mutex
}

func New(path string) *JSONStore {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.WriteFile(path, []byte("[]"), 0644)
	}
	return &JSONStore{
		filePath: path,
	}
}

func (s *JSONStore) LoadTasks() ([]entity.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.Open(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []entity.Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []entity.Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *JSONStore) SaveTasks(tasks []entity.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.Create(s.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tasks)
}
