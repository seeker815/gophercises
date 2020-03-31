package cmd

import (
	//homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a CLI for managing your TODOs.",
	Long:  `task is a CLI for managing your TODOs.`,
	
	Run: func(cmd *cobra.Command, args []string) { },
}
