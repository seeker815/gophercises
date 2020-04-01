package cmd

import (
	"fmt"
	"strconv"

	"github.com/seeker815/gophercises/task/db"
	"github.com/spf13/cobra"
)

// rmCmd removes a unwanted task instead of completing it
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "rm deletes a task instead of completing it",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		// ensure user invokes the taskId
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the input for taskID ")
				return
			}
			ids = append(ids, id)
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}

		// Delete the task from bucket
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to delete \"%d\" task from the list.\n", id)
			}
			fmt.Printf("Deleted task \"%d\" from the todo list.\n", id)
		}
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
