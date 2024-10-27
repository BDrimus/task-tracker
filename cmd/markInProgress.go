/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/BDrimus/task-tracker/task"
	"github.com/spf13/cobra"
)

// markInProgressCmd represents the markInProgress command
var markInProgressCmd = &cobra.Command{
	Use:   "markInProgress [task-id]",
	Short: "Mark task as in progress",
	Long: `Mark task as in progress using its ID. For example:

markInProgress 3`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		idStr := args[0]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			fmt.Println("ID has to be a positive number")
		}

		task.UpdateTask(id, task.Task{Status: task.InProgress})

		fmt.Println("markInProgress called")
	},
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markInProgressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markInProgressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
