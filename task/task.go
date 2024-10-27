package task

import (
	"encoding/json"
	"errors"
	"os"
	"time"

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
	Id          uint64         `json:"id"`
	Description string         `json:"description"`
	Status      ProgressStatus `json:"status"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

func Initialise() {
	if _, err := os.Stat(DBLocation); errors.Is(err, os.ErrNotExist) {

		err := os.WriteFile(DBLocation, []byte(""), 0644)
		if err != nil {
			slog.Error("unable to create file: %w", err)
		}
	}
}

func writeToJson(tasks []Task) error {
	marshalled, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		slog.Error("Couldn't marshal")
		return err
	}

	_ = os.WriteFile(DBLocation, marshalled, 0644)

	return nil
}

// getNextAvailableId retrieves the next available task ID.
// It fetches the list of tasks using the GetTasks function.
// If the task list is empty or the JSON is empty, it returns 1 as the next available ID.
// Otherwise, it returns the ID of the last task incremented by 1.
// Returns an error if there is an issue fetching the tasks.
func getNextAvailableId() (uint64, error) {

	tasks, err := GetTasks()
	if err != nil {
		if err.Error() == ErrEmptyJson.Error() {
			return 1, nil
		}
		return 0, err
	}

	if len(tasks) == 0 {
		return 1, nil
	}

	lastTask := tasks[len(tasks)-1]

	return lastTask.Id + 1, nil
}

// AddTask adds a new task with the given description to the task list.
//
// Parameters:
//
//	description (string): The description of the new task.
//
// Returns:
//
//	*Task: A pointer to the newly created Task.
//	error: An error if the task could not be added, or nil if the task was added successfully.
func AddTask(description string) (*Task, error) {

	nextAvailableId, err := getNextAvailableId()
	if err != nil {
		slog.Error("Couldn't find available ID")
		return nil, err
	}

	t := Task{
		Id:          nextAvailableId,
		Description: description,
		Status:      NotStarted,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks, err := GetTasks()
	if err != nil && err.Error() != ErrEmptyJson.Error() {
		slog.Error(ErrCouldntGetTasks.Error())
		return nil, err
	}

	tasks = append(tasks, t)

	writeToJson(tasks)

	return &t, nil
}

func UpdateTask(id uint64, description string) (*Task, error) {
	tasks, err := GetTasks()
	if err != nil {
		return nil, ErrCouldntGetTasks
	}

	var taskToUpdate Task
	var listOfTasks []Task

	for _, task := range tasks {
		if task.Id == id {
			taskToUpdate = task
			taskToUpdate.Description = description
			listOfTasks = append(listOfTasks, taskToUpdate)

			continue
		}
		listOfTasks = append(listOfTasks, task)
	}

	writeToJson(listOfTasks)

	return &taskToUpdate, nil
}

func DeleteTask(id uint64) error {
	tasks, err := GetTasks()
	if err != nil {
		slog.Error(ErrCouldntGetTasks.Error())
		return err
	}

	var listOfTasks []Task

	for _, task := range tasks {

		if task.Id == id {
			continue
		}

		listOfTasks = append(listOfTasks, task)

		writeToJson(listOfTasks)
	}

	return nil
}

// FIXME - Filter options

// GetTasks retrieves the list of tasks from the database file specified by DBLocation.
// It reads the file content, unmarshals the JSON data into a slice of Task structs, and returns it.
// If the file is empty, it returns an ErrEmptyJson error.
// If there is an error reading the file or unmarshalling the JSON data, it returns the respective error.
//
// Returns:
//   - []Task: A slice of Task structs representing the tasks.
//   - error: An error if there is an issue reading the file or unmarshalling the JSON data, or if the file is empty.
func GetTasks() ([]Task, error) {

	var t []Task

	f, err := os.ReadFile(DBLocation)
	if err != nil {
		return nil, err
	}

	if len(f) == 0 {
		return nil, ErrEmptyJson
	}

	err = json.Unmarshal(f, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

var ErrEmptyJson = errors.New("empty json")
var ErrCouldntGetTasks = errors.New("couln't get tasks")
