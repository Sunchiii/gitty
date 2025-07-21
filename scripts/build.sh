#!/bin/bash

# Build script for Gitty
set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Default values
VERSION=${VERSION:-"v0.1.0"}
BUILD_DATE=$(date -u +"%Y-%m-%d")
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "development")
LDFLAGS="-X github.com/Sunchiii/gitty/cmd.Version=$VERSION -X github.com/Sunchiii/gitty/cmd.BuildDate=$BUILD_DATE -X github.com/Sunchiii/gitty/cmd.GitCommit=$GIT_COMMIT"

echo -e "${GREEN}Building Gitty...${NC}"
echo "Version: $VERSION"
echo "Build Date: $BUILD_DATE"
echo "Git Commit: $GIT_COMMIT"

# Update version.txt
echo "$VERSION" > cmd/version.txt

# Build for current platform
echo -e "${YELLOW}Building for $(go env GOOS)/$(go env GOARCH)...${NC}"
go build -ldflags "$LDFLAGS" -o gitty

# Test the build
echo -e "${YELLOW}Testing build...${NC}"
./gitty version

echo -e "${GREEN}Build completed successfully!${NC}"
echo "Binary: ./gitty" 