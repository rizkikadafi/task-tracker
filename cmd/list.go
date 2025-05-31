/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/rizkikadafi/task-tracker/core"
	"github.com/rizkikadafi/task-tracker/entity"
	"github.com/rizkikadafi/task-tracker/store/jsonstore"
	"github.com/spf13/cobra"
)

var statusFilter string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// load store
		storePath := "tasks.json"
		store := jsonstore.New(storePath)

		// create service
		service, err := core.NewStoreService(store)
		if err != nil {
			fmt.Println("Error initializing service:", err)
			os.Exit(1)
		}

		var tasks []entity.Task
		if statusFilter != "" {
			status, err := entity.ParseStatus(statusFilter)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

			tasks, err = service.ListTasksByStatus(status)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
		} else {
			tasks = service.ListTasks()
		}

		if len(tasks) == 0 {
			fmt.Printf("No tasks found for status (%s).\n", statusFilter)
			return
		}

		fmt.Println("Tasks:")
		for _, task := range tasks {
			fmt.Printf("[%d] %s (%s)\n", task.ID, task.Title, task.Status)
		}
	},
}

func init() {
	listCmd.Flags().StringVarP(&statusFilter, "status", "s", "", "Filter by status (todo, in-progress, done)")
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func isValidStatus(s string) bool {
	switch s {
	case "todo", "in-progress", "done":
		return true
	default:
		return false
	}
}
