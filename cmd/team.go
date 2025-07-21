package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "Show team workflow information",
	Run: func(cmd *cobra.Command, args []string) {
		config := LoadConfig()
		fmt.Println("🏢 Team Workflow Guide:")
		fmt.Println("")
		fmt.Println("📋 Branch Strategy:")
		for _, branch := range config.ProtectedBranches {
			fmt.Printf("  • %s - Protected branch\n", branch)
		}
		fmt.Println("  • feature/* - New features")
		fmt.Println("  • hotfix/* - Production fixes")
		fmt.Println("")
		fmt.Println("🔄 Workflow:")
		fmt.Println("  1. gitty start feature <name>")
		fmt.Println("  2. Make changes and commit")
		fmt.Println("  3. gitty sync (rebase with develop)")
		fmt.Println("  4. git push origin feature/<name>")
		fmt.Println("  5. Create pull request")
		fmt.Println("  6. After merge: gitty finish feature <name>")
		fmt.Println("")
		fmt.Println("🛡️  Protection:")
		fmt.Println("  • Never commit directly to main/uat/develop")
		fmt.Println("  • Always sync before pushing")
		fmt.Println("  • Use descriptive commit messages")
		fmt.Println("")
		fmt.Println("🔧 Useful Commands:")
		fmt.Println("  • gitty status    - Check current status")
		fmt.Println("  • gitty validate  - Validate branch naming")
		fmt.Println("  • gitty conflict  - Check for conflicts")
		fmt.Println("  • gitty cleanup   - Clean up merged branches")
	},
}
