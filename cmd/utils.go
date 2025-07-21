package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Helper function to execute git commands
func executeCommands(cmds [][]string) {
	for _, cmd := range cmds {
		fmt.Printf("Executing: %s\n", strings.Join(cmd, " "))
		command := exec.Command(cmd[0], cmd[1:]...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		if err := command.Run(); err != nil {
			log.Fatalf("Command failed: %v", err)
		}
	}
}

// Helper function to execute git commands and return output
func executeCommandWithOutput(cmd []string) string {
	command := exec.Command(cmd[0], cmd[1:]...)
	output, err := command.Output()
	if err != nil {
		log.Fatalf("Command failed: %v", err)
	}
	return strings.TrimSpace(string(output))
}

// Helper function to get current branch name
func getCurrentBranch() string {
	return executeCommandWithOutput([]string{"git", "rev-parse", "--abbrev-ref", "HEAD"})
}

// Helper function to write git hook
func writeHook(path, content string) error {
	// Ensure .git/hooks directory exists
	dir := ".git/hooks"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(path, []byte(content), 0755)
}

// Helper function to check if branch exists
func branchExists(branch string) bool {
	cmd := exec.Command("git", "show-ref", "--verify", "--quiet", "refs/heads/"+branch)
	return cmd.Run() == nil
}
