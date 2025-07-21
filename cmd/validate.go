package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate current branch follows naming conventions",
	Run: func(cmd *cobra.Command, args []string) {
		branch := getCurrentBranch()
		fmt.Printf("Validating branch: %s\n", branch)

		config := LoadConfig()
		// Check if it's a protected branch
		if config.IsProtectedBranch(branch) {
			fmt.Printf("✅ %s is a protected branch\n", branch)
			return
		}

		// Check naming convention
		if strings.HasPrefix(branch, "feature/") || strings.HasPrefix(branch, "hotfix/") {
			fmt.Printf("✅ %s follows naming convention\n", branch)
		} else {
			fmt.Printf("⚠️  %s doesn't follow naming convention (should be feature/ or hotfix/)\n", branch)
		}

		// Check if branch is up to date
		executeCommands([][]string{
			{"git", "fetch", "origin"},
		})

		behind := executeCommandWithOutput([]string{"git", "rev-list", "--count", "origin/develop..HEAD"})
		if behind == "0" {
			fmt.Println("✅ Branch is up to date with develop")
		} else {
			fmt.Printf("⚠️  Branch is %s commits behind develop\n", behind)
		}
	},
}
