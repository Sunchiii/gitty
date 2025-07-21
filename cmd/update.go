package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// ReleaseInfo represents GitHub release information
type ReleaseInfo struct {
	TagName     string  `json:"tag_name"`
	Body        string  `json:"body"`
	Assets      []Asset `json:"assets"`
	PublishedAt string  `json:"published_at"`
}

type Asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check for updates and update gitty to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔍 Checking for updates...")

		// Get current version
		currentVersion := getCurrentVersion()
		fmt.Printf("Current version: %s\n", currentVersion)

		// Check for latest version
		latestRelease, err := getLatestRelease()
		if err != nil {
			fmt.Printf("❌ Failed to check for updates: %v\n", err)
			return
		}

		fmt.Printf("Latest version: %s\n", latestRelease.TagName)

		if isNewerVersion(currentVersion, latestRelease.TagName) {
			fmt.Println("🎉 A new version is available!")
			fmt.Printf("Release notes:\n%s\n", latestRelease.Body)

			// Ask user if they want to update
			fmt.Print("Do you want to update? (y/N): ")
			var response string
			fmt.Scanln(&response)

			if strings.ToLower(response) == "y" || strings.ToLower(response) == "yes" {
				if err := performUpdate(latestRelease); err != nil {
					fmt.Printf("❌ Update failed: %v\n", err)
					return
				}
				fmt.Println("✅ Update completed successfully!")
			} else {
				fmt.Println("Update cancelled.")
			}
		} else {
			fmt.Println("✅ You're already running the latest version!")
		}
	},
}

// getLatestRelease fetches the latest release from GitHub
func getLatestRelease() (*ReleaseInfo, error) {
	// Replace with your actual GitHub repository URL
	url := "https://api.github.com/repos/Sunchiii/gitty/releases/latest"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var release ReleaseInfo
	if err := json.Unmarshal(body, &release); err != nil {
		return nil, err
	}

	return &release, nil
}

// isNewerVersion compares two version strings
func isNewerVersion(current, latest string) bool {
	// Simple version comparison - in production you'd want a more robust version parser
	return latest > current
}

// performUpdate downloads and installs the latest version
func performUpdate(release *ReleaseInfo) error {
	fmt.Println("📥 Downloading update...")

	// Find the appropriate asset for the current platform
	var downloadURL string
	platform := runtime.GOOS
	arch := runtime.GOARCH

	for _, asset := range release.Assets {
		if strings.Contains(asset.Name, platform) && strings.Contains(asset.Name, arch) {
			downloadURL = asset.BrowserDownloadURL
			break
		}
	}

	if downloadURL == "" {
		return fmt.Errorf("no compatible release found for %s/%s", platform, arch)
	}

	// Download the new version
	resp, err := http.Get(downloadURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create temporary file
	tempFile, err := os.CreateTemp("", "gitty-update-*")
	if err != nil {
		return err
	}
	defer os.Remove(tempFile.Name())

	// Write the downloaded file
	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		return err
	}
	tempFile.Close()

	// Make the file executable
	if err := os.Chmod(tempFile.Name(), 0755); err != nil {
		return err
	}

	// Get the current executable path
	currentPath, err := os.Executable()
	if err != nil {
		return err
	}

	// Create backup
	backupPath := currentPath + ".backup"
	if err := os.Rename(currentPath, backupPath); err != nil {
		return err
	}

	// Move new version to current location
	if err := os.Rename(tempFile.Name(), currentPath); err != nil {
		// Restore backup on failure
		os.Rename(backupPath, currentPath)
		return err
	}

	// Remove backup
	os.Remove(backupPath)

	return nil
}

// Add a version command to show current version
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show gitty version information",
	Run: func(cmd *cobra.Command, args []string) {
		versionInfo := GetVersionInfo()
		fmt.Printf("Gitty version: %s\n", versionInfo["version"])
		fmt.Printf("Build date: %s\n", versionInfo["buildDate"])
		fmt.Printf("Git commit: %s\n", versionInfo["gitCommit"])
		fmt.Printf("Go version: %s\n", versionInfo["goVersion"])
		fmt.Printf("Platform: %s\n", versionInfo["platform"])
		fmt.Printf("Compiler: %s\n", versionInfo["compiler"])
	},
}
