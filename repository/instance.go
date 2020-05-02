package repository

import (
	"fmt"

	"github.com/sato11/vscode-remote-development-go/model/tasks"
)

type instance struct {
	tasks []tasks.Task
}

// New initializes instance
func New() tasks.Repository {
	s := new(instance)
	s.tasks = make([]tasks.Task, 2, 20)
	s.tasks[0] = tasks.Task{
		ID:   1,
		Text: "task1",
		Done: false,
	}
	s.tasks[1] = tasks.Task{
		ID:   2,
		Text: "task2",
		Done: false,
	}
	return s
}

func (s *instance) Add(task tasks.Task) int {
	task.ID = len(s.tasks) + 1
	s.tasks = append(s.tasks, task)
	return task.ID
}

func (s *instance) List() []*tasks.Task {
	result := []*tasks.Task{}
	for i, task := range s.tasks {
		if !task.Done {
			result = append(result, &s.tasks[i])
		}
	}
	return result
}

func (s *instance) Done(id int) error {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("Record Not Found. ID: %d", id)
}
