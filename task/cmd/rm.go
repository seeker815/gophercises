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

		if len(ids) == 0 {
			fmt.Println("No available tasks in the list to remove")
			return
		}
		// Delete the task from bucket
		for _, id := range ids {
			if id <= 0 {
				fmt.Println("Invalid task number:", id)
				continue
			}
			err := db.DeleteTask(id)
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
