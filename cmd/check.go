package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check branch rules before push",
	Run: func(cmd *cobra.Command, args []string) {
		config := LoadConfig()
		branch := getCurrentBranch()
		if config.IsProtectedBranch(branch) {
			log.Fatalf("❌ Do not commit or push directly to %s", branch)
		}

		fmt.Printf("Checking branch: %s\n", branch)

		// Check for uncommitted changes
		status := executeCommandWithOutput([]string{"git", "status", "--porcelain"})
		if status != "" {
			fmt.Println("⚠️  Warning: You have uncommitted changes")
		}

		executeCommands([][]string{
			{"git", "fetch", "origin"},
			{"git", "status"},
			{"git", "log", "origin/develop..HEAD"},
		})

		fmt.Printf("✅ Branch %s is ready for push\n", branch)
	},
}
