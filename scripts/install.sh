#!/bin/bash

# Gitty Installation Script
set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
REPO="Sunchiii/gitty"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="gitty"

# Get the latest version
get_latest_version() {
    curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/'
}

# Get the appropriate asset URL for the current platform
get_asset_url() {
    local version=$1
    local platform=$(uname -s | tr '[:upper:]' '[:lower:]')
    local arch=$(uname -m)
    
    # Map architecture names
    case $arch in
        x86_64) arch="amd64" ;;
        arm64) arch="arm64" ;;
        aarch64) arch="arm64" ;;
    esac
    
    # Map platform names
    case $platform in
        darwin) platform="darwin" ;;
        linux) platform="linux" ;;
        msys*|cygwin*|mingw*) platform="windows" ;;
        *) echo "Unsupported platform: $platform" >&2; exit 1 ;;
    esac
    
    local asset_name="gitty-$version-$platform-$arch"
    if [ "$platform" = "windows" ]; then
        asset_name="$asset_name.exe"
    fi
    
    echo "https://github.com/$REPO/releases/download/$version/$asset_name"
}

# Download and install
install_gitty() {
    local version=$1
    local download_url=$2
    
    echo -e "${BLUE}Downloading Gitty $version...${NC}"
    
    # Create temporary directory
    local temp_dir=$(mktemp -d)
    local temp_file="$temp_dir/gitty"
    
    # Download the binary
    if ! curl -L -o "$temp_file" "$download_url"; then
        echo -e "${RED}Failed to download Gitty${NC}"
        rm -rf "$temp_dir"
        exit 1
    fi
    
    # Make it executable
    chmod +x "$temp_file"
    
    # Check if we can write to install directory
    if [ ! -w "$INSTALL_DIR" ]; then
        echo -e "${YELLOW}Need sudo permissions to install to $INSTALL_DIR${NC}"
        if ! sudo mv "$temp_file" "$INSTALL_DIR/$BINARY_NAME"; then
            echo -e "${RED}Failed to install Gitty${NC}"
            rm -rf "$temp_dir"
            exit 1
        fi
    else
        mv "$temp_file" "$INSTALL_DIR/$BINARY_NAME"
    fi
    
    # Clean up
    rm -rf "$temp_dir"
    
    echo -e "${GREEN}Gitty $version installed successfully!${NC}"
}

# Verify installation
verify_installation() {
    if command -v "$BINARY_NAME" >/dev/null 2>&1; then
        echo -e "${GREEN}Installation verified!${NC}"
        echo -e "${BLUE}Run 'gitty --help' to get started${NC}"
    else
        echo -e "${RED}Installation failed${NC}"
        exit 1
    fi
}

# Main installation function
main() {
    echo -e "${BLUE}Installing Gitty...${NC}"
    
    # Check if gitty is already installed
    if command -v "$BINARY_NAME" >/dev/null 2>&1; then
        echo -e "${YELLOW}Gitty is already installed.${NC}"
        read -p "Do you want to update it? (y/N): " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            echo "Installation cancelled."
            exit 0
        fi
    fi
    
    # Get latest version
    echo -e "${BLUE}Fetching latest version...${NC}"
    local latest_version=$(get_latest_version)
    
    if [ -z "$latest_version" ]; then
        echo -e "${RED}Failed to get latest version${NC}"
        exit 1
    fi
    
    echo -e "${GREEN}Latest version: $latest_version${NC}"
    
    # Get download URL
    local download_url=$(get_asset_url "$latest_version")
    
    if [ -z "$download_url" ]; then
        echo -e "${RED}Failed to get download URL${NC}"
        exit 1
    fi
    
    # Install
    install_gitty "$latest_version" "$download_url"
    
    # Verify
    verify_installation
}

# Handle command line arguments
case "${1:-}" in
    --version)
        get_latest_version
        ;;
    --help)
        echo "Gitty Installation Script"
        echo ""
        echo "Usage:"
        echo "  curl -fsSL https://raw.githubusercontent.com/$REPO/main/scripts/install.sh | bash"
        echo "  curl -fsSL https://raw.githubusercontent.com/$REPO/main/scripts/install.sh | bash -s -- --version"
        echo ""
        echo "Options:"
        echo "  --version    Show latest version"
        echo "  --help       Show this help"
        ;;
    *)
        main
        ;;
esac 