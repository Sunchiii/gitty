package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current branch status and pending changes",
	Run: func(cmd *cobra.Command, args []string) {
		branch := getCurrentBranch()
		fmt.Printf("Current branch: %s\n", branch)

		executeCommands([][]string{
			{"git", "fetch", "origin"},
			{"git", "status"},
			{"git", "log", "--oneline", "-5"},
		})
	},
}
