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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [task-id] [task-description]",
	Short: "Update an existing task",
	Long: `Update an existing task using the ID and a new description. For example:

update 3 "this is an updated task"`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		idStr, description := args[0], args[1]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			fmt.Println("ID has to be a positive number")
		} else {
			task.UpdateTask(id, description)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
