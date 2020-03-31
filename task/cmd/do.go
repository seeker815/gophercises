package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as completed",
	Long:  `Marks a task as completed in the todo list`,
	Run: func(cmd *cobra.Command, args []string) {
		//call do with numbers rather than task name
		var taskIDs []int
		for _, arg := range args {
			taskID, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Please enter task item number to mark task as done")
				return
			}

			if taskID > 0 {
				taskIDs = append(taskIDs, taskID)
			} else {
				fmt.Println("Enter a valid task item number")
			}
		}
		fmt.Println("Tasks marked as complete ", taskIDs)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
