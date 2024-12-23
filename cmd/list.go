/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"

	"github.com/BDrimus/task-tracker/task"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var tasks []task.Task
		var err error

		if len(args) == 1 {
			switch args[0] {
			case "todo":
				tasks, err = task.GetTasks(task.NotStarted)
				if err != nil {
					slog.Error(err.Error())
				}
			case "inProgress":
				tasks, err = task.GetTasks(task.InProgress)
				if err != nil {
					slog.Error(err.Error())
				}
			case "done":
				tasks, err = task.GetTasks(task.Done)
				if err != nil {
					slog.Error(err.Error())
				}
			default:
				fmt.Println("Invalid task status")
			}
		} else {
			tasks, err = task.GetTasks()
			if err != nil {
				slog.Error(err.Error())
			}
		}

		for _, task := range tasks {
			fmt.Println(task)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
