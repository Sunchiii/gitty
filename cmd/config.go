package cmd

import (
	"os"
	"strings"
)

// Config holds the application configuration
type Config struct {
	ProtectedBranches []string
	DefaultBaseBranch string
	HookTemplate      string
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		ProtectedBranches: []string{"main", "uat", "develop"},
		DefaultBaseBranch: "develop",
		HookTemplate: `#!/bin/sh
branch=$(git symbolic-ref HEAD | sed -e 's,.*/,,')
protected_branches="main uat develop"
for protected in $protected_branches; do
  if [ "$branch" = "$protected" ]; then
    echo "[GITTY] Push to $branch is not allowed!"
    exit 1
  fi
done`,
	}
}

// IsProtectedBranch checks if a branch is protected
func (c *Config) IsProtectedBranch(branch string) bool {
	for _, protected := range c.ProtectedBranches {
		if branch == protected {
			return true
		}
	}
	return false
}

// GetBaseBranch returns the appropriate base branch for a branch type
func (c *Config) GetBaseBranch(branchType string) string {
	if branchType == "hotfix" {
		return "main"
	}
	return c.DefaultBaseBranch
}

// GetTargetBranch returns the appropriate target branch for a branch type
func (c *Config) GetTargetBranch(branchType string) string {
	if branchType == "hotfix" {
		return "main"
	}
	return c.DefaultBaseBranch
}

// LoadConfig loads configuration from environment or uses defaults
func LoadConfig() *Config {
	config := DefaultConfig()

	// Allow customization via environment variables
	if protected := os.Getenv("GITTY_PROTECTED_BRANCHES"); protected != "" {
		config.ProtectedBranches = strings.Split(protected, ",")
	}

	if base := os.Getenv("GITTY_DEFAULT_BASE"); base != "" {
		config.DefaultBaseBranch = base
	}

	return config
}
