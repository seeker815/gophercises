package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rmCmd removes a unwanted task instead of completing it
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "rm deletes a task instead of completing it",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rm called")
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
