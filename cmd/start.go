package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [feature|hotfix] <name>",
	Short: "Start a new feature or hotfix branch",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		type_, name := args[0], args[1]

		// Validate branch type
		if type_ != "feature" && type_ != "hotfix" {
			log.Fatal("Branch type must be 'feature' or 'hotfix'")
		}

		// Validate branch name format
		if strings.Contains(name, " ") {
			log.Fatal("Branch name cannot contain spaces. Use hyphens or underscores.")
		}

		config := LoadConfig()
		base := config.GetBaseBranch(type_)

		branchName := type_ + "/" + name

		// Check if branch already exists
		if branchExists(branchName) {
			log.Fatalf("Branch %s already exists", branchName)
		}

		fmt.Printf("Starting %s branch: %s\n", type_, name)

		cmds := [][]string{
			{"git", "checkout", base},
			{"git", "pull", "origin", base},
			{"git", "checkout", "-b", branchName},
		}
		executeCommands(cmds)

		fmt.Printf("✅ Successfully created %s branch: %s\n", type_, branchName)
		fmt.Printf("💡 Next steps:\n")
		fmt.Printf("   1. Make your changes\n")
		fmt.Printf("   2. git add . && git commit -m \"your message\"\n")
		fmt.Printf("   3. gitty sync (to sync with develop)\n")
		fmt.Printf("   4. git push origin %s\n", branchName)
	},
}
