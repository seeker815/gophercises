package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the todo",
	Long:  `Adds a task to the todo`,
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to the task list ", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
