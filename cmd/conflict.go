package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var conflictCmd = &cobra.Command{
	Use:   "conflict",
	Short: "Check for potential merge conflicts",
	Run: func(cmd *cobra.Command, args []string) {
		branch := getCurrentBranch()
		fmt.Printf("Checking for conflicts in: %s\n", branch)

		executeCommands([][]string{
			{"git", "fetch", "origin"},
		})

		// Check if there would be conflicts with develop
		mergeCmd := exec.Command("git", "merge-tree", "origin/develop", "HEAD")
		_, err := mergeCmd.Output()
		if err != nil {
			fmt.Println("⚠️  Potential conflicts detected")
		} else {
			fmt.Println("✅ No conflicts detected")
		}
	},
}
