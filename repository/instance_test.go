package repository

import (
	"testing"

	"github.com/sato11/vscode-remote-development-go/model/tasks"
)

func TestAdd(t *testing.T) {
	t.Run("Successfully inserts", func(t *testing.T) {
		rep := New()
		rep.Add(tasks.Task{
			Text: "new task",
		})

		if len(rep.(*instance).tasks) != 3 {
			t.Errorf("There should be 3 tasks, got: %d", len(rep.(*instance).tasks))
		}

		addedTask := rep.(*instance).tasks[2]
		if addedTask.Text != "new task" {
			t.Errorf("Text of the new task should be 'new task', got: %s", addedTask.Text)
		}
		if addedTask.ID <= 2 {
			t.Errorf("New task should be assigned a new ID, got: %d", addedTask.ID)
		}

		for i, task := range rep.(*instance).tasks {
			if i != 2 && addedTask.ID == task.ID {
				t.Errorf("New task should be assigned a unique ID: got duplicated %d", addedTask.ID)
			}
			if i != 2 && addedTask.Text == task.Text {
				t.Errorf("New task should not affect existing task: got duplicated %s", addedTask.Text)
			}
		}
	})
}

func TestList(t *testing.T) {
	t.Run("Successfully returns list", func(t *testing.T) {
		rep := New()
		tasks := rep.List()
		if len(tasks) != 2 {
			t.Errorf("Should return 2 tasks, got: %d", len(tasks))
		}

		for _, task := range tasks {
			if task.Done {
				t.Errorf("Task marked as Done should not be included, id: %d", task.ID)
			}
		}
	})
}

func TestDone(t *testing.T) {
	t.Run("Successfully marks as done when valid ID is provided", func(t *testing.T) {
		id := 1

		rep := New()
		err := rep.Done(id)

		if err != nil {
			t.Errorf("Should be successful, got: %s", err)
		}
	})

	t.Run("Should return error when invalid ID is provided", func(t *testing.T) {
		id := 1 << 10

		rep := New()
		err := rep.Done(id)

		if err == nil {
			t.Errorf("Record should not exist, id: %d", id)
		}
	})
}
