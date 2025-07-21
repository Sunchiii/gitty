#!/bin/bash

# Gitty Uninstall Script
set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
BINARY_NAME="gitty"
INSTALL_DIR="/usr/local/bin"
BINARY_PATH="$INSTALL_DIR/$BINARY_NAME"

echo -e "${BLUE}Uninstalling Gitty...${NC}"

# Check if gitty is installed
if [ ! -f "$BINARY_PATH" ]; then
    echo -e "${YELLOW}Gitty is not installed at $BINARY_PATH${NC}"
    exit 0
fi

# Show current version
if command -v "$BINARY_NAME" >/dev/null 2>&1; then
    echo -e "${BLUE}Current version:${NC}"
    "$BINARY_NAME" version
    echo
fi

# Ask for confirmation
read -p "Are you sure you want to uninstall Gitty? (y/N): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Uninstall cancelled."
    exit 0
fi

# Remove the binary
echo -e "${BLUE}Removing Gitty binary...${NC}"
if [ -w "$INSTALL_DIR" ]; then
    rm -f "$BINARY_PATH"
else
    echo -e "${YELLOW}Need sudo permissions to remove from $INSTALL_DIR${NC}"
    sudo rm -f "$BINARY_PATH"
fi

# Verify removal
if [ ! -f "$BINARY_PATH" ]; then
    echo -e "${GREEN}Gitty uninstalled successfully!${NC}"
else
    echo -e "${RED}Failed to uninstall Gitty${NC}"
    exit 1
fi

echo -e "${BLUE}Thanks for using Gitty!${NC}" 