package main

import (
	"fmt"

	"github.com/BDrimus/task-tracker/task"
)

func main() {
	fmt.Println("Test")

	tasks, _ := task.GetTasks()
	// task.AddTask("TestTask")

	fmt.Println(tasks)
}
