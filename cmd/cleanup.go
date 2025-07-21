package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Clean up merged branches and sync with remote",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cleaning up merged branches...")

		executeCommands([][]string{
			{"git", "fetch", "--prune"},
		})

		// Get merged branches
		mergedDevelop := executeCommandWithOutput([]string{"git", "branch", "--merged", "develop"})
		mergedMain := executeCommandWithOutput([]string{"git", "branch", "--merged", "main"})

		fmt.Println("Merged branches that can be deleted:")
		fmt.Println(mergedDevelop)
		fmt.Println(mergedMain)

		fmt.Println("✅ Cleanup completed")
	},
}
