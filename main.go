/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/BDrimus/task-tracker/cmd"
	"github.com/BDrimus/task-tracker/task"
)

func main() {
	task.Initialise()
	cmd.Execute()
}
