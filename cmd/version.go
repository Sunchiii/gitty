package cmd

import (
	"embed"
	"runtime"
	"strings"
)

//go:embed version.txt
var versionFile embed.FS

// Version information
const (
	Version   = "v0.1.0"
	BuildDate = "2024-01-01"
	GitCommit = "development"
)

// getCurrentVersion returns the current version of gitty
func getCurrentVersion() string {
	// Try to read from embedded version file first
	if content, err := versionFile.ReadFile("version.txt"); err == nil {
		return strings.TrimSpace(string(content))
	}

	// Fallback to constant
	return Version
}

// GetVersionInfo returns detailed version information
func GetVersionInfo() map[string]string {
	return map[string]string{
		"version":   getCurrentVersion(),
		"buildDate": BuildDate,
		"gitCommit": GitCommit,
		"goVersion": runtime.Version(),
		"platform":  runtime.GOOS + "/" + runtime.GOARCH,
		"compiler":  runtime.Compiler,
	}
}
