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

// markDoneCmd represents the markDone command
var markDoneCmd = &cobra.Command{
	Use:   "markDone [task-id]",
	Short: "Mark task as done",
	Long: `Mark task as done using its ID. For example:

markDone 3`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		idStr := args[0]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			fmt.Println("ID has to be a positive number")
		}

		task.UpdateTask(id, task.Task{Status: task.Done})
	},
}

func init() {
	rootCmd.AddCommand(markDoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markDoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markDoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
