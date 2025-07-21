package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync current branch with develop",
	Run: func(cmd *cobra.Command, args []string) {
		currentBranch := getCurrentBranch()
		fmt.Printf("Syncing branch: %s\n", currentBranch)

		// Check for uncommitted changes
		status := executeCommandWithOutput([]string{"git", "status", "--porcelain"})
		if status != "" {
			log.Fatal("❌ You have uncommitted changes. Please commit or stash them first.")
		}

		executeCommands([][]string{
			{"git", "fetch", "origin"},
			{"git", "rebase", "origin/develop"},
		})

		fmt.Printf("✅ Successfully synced %s with develop\n", currentBranch)
	},
}
