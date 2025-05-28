/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/rizkikadafi/task-tracker/core"
	"github.com/rizkikadafi/task-tracker/store/jsonstore"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task title]",
	Short: "Add new task",
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]

		// load store
		storePath := "tasks.json"
		store := jsonstore.New(storePath)

		// create service
		service, err := core.NewStoreService(store)
		if err != nil {
			fmt.Println("Error initializing service:", err)
			os.Exit(1)
		}

		// add task
		task, err := service.AddTask(title)
		if err != nil {
			fmt.Println("Failed to add task:", err)
			os.Exit(1)
		}

		fmt.Printf("Task added: [%d] %s (%s)\n", task.ID, task.Title, task.Status)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
