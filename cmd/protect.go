package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var protectCmd = &cobra.Command{
	Use:   "protect",
	Short: "Install Git hook to protect critical branches",
	Run: func(cmd *cobra.Command, args []string) {
		config := LoadConfig()
		hook := config.HookTemplate
		err := writeHook(".git/hooks/pre-push", hook)
		if err != nil {
			log.Fatal("Failed to write hook:", err)
		}
		log.Println("✅ Hook installed successfully.")
	},
}
