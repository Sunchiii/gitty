package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var finishCmd = &cobra.Command{
	Use:   "finish [feature|hotfix] <name>",
	Short: "Finish a feature or hotfix branch",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		type_, name := args[0], args[1]

		// Validate branch type
		if type_ != "feature" && type_ != "hotfix" {
			log.Fatal("Branch type must be 'feature' or 'hotfix'")
		}

		config := LoadConfig()
		target := config.GetTargetBranch(type_)
		branch := type_ + "/" + name

		// Check if branch exists
		if !branchExists(branch) {
			log.Fatalf("Branch %s does not exist", branch)
		}

		fmt.Printf("Finishing %s branch: %s\n", type_, branch)

		// Check for uncommitted changes
		status := executeCommandWithOutput([]string{"git", "status", "--porcelain"})
		if status != "" {
			log.Fatal("❌ You have uncommitted changes. Please commit or stash them first.")
		}

		cmds := [][]string{
			{"git", "checkout", target},
			{"git", "pull", "origin", target},
			{"git", "merge", "--no-ff", branch},
			{"git", "branch", "-d", branch},
		}
		executeCommands(cmds)

		fmt.Printf("✅ Successfully finished %s branch: %s\n", type_, branch)
		fmt.Printf("💡 Don't forget to push: git push origin %s\n", target)
	},
}
