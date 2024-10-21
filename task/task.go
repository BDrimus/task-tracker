package task

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"golang.org/x/exp/slog"
)

const DBLocation = "tasks.json"

type ProgressStatus string

const (
	NotStarted ProgressStatus = "Not Started"
	InProgress ProgressStatus = "In Progress"
	Done       ProgressStatus = "Done"
)

type Task struct {
	// ID     uint64         `json:"id"`
	Title  string         `json:"title"`
	Status ProgressStatus `json:"status"`
}

func Initialise() {
	if _, err := os.Stat(DBLocation); errors.Is(err, os.ErrNotExist) {

		err := os.WriteFile(DBLocation, []byte(""), 0644)
		if err != nil {
			slog.Error("unable to create file: %w", err)
		}
	}
}

func FindSmallestAvailableID() (id uint64) {

	return id
}

// FIXME - Overwrite
func AddTask(title string) ([]Task, error) {

	var t []Task

	t = append(t, Task{Title: title, Status: NotStarted})

	marshalled, _ := json.MarshalIndent(t, "", "  ")

	_ = os.WriteFile(DBLocation, marshalled, 0644)

	return t, nil
}

func GetTasks() ([]Task, error) {

	var t []Task

	f, err := os.ReadFile(DBLocation)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(f, &t)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return t, nil
}
