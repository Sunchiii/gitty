#!/bin/bash

# Release script for Gitty
set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if version is provided
if [ -z "$1" ]; then
    echo -e "${RED}Usage: $0 <version>${NC}"
    echo "Example: $0 v1.0.0"
    exit 1
fi

VERSION=$1
BUILD_DATE=$(date -u +"%Y-%m-%d")
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "development")

echo -e "${GREEN}Creating release for version: $VERSION${NC}"

# Create release directory
RELEASE_DIR="release/$VERSION"
mkdir -p "$RELEASE_DIR"

# Update version.txt
echo "$VERSION" > cmd/version.txt

# Build for multiple platforms
PLATFORMS=(
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
)

for platform in "${PLATFORMS[@]}"; do
    IFS='/' read -r GOOS GOARCH <<< "$platform"
    
    echo -e "${YELLOW}Building for $GOOS/$GOARCH...${NC}"
    
    # Set build flags
    LDFLAGS="-X github.com/Sunchiii/gitty/cmd.Version=$VERSION -X github.com/Sunchiii/gitty/cmd.BuildDate=$BUILD_DATE -X github.com/Sunchiii/gitty/cmd.GitCommit=$GIT_COMMIT"
    
    # Set output name
    OUTPUT="gitty-$VERSION-$GOOS-$GOARCH"
    if [ "$GOOS" = "windows" ]; then
        OUTPUT="$OUTPUT.exe"
    fi
    
    # Build
    GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "$LDFLAGS" -o "$RELEASE_DIR/$OUTPUT"
    
    # Create checksum
    if [ "$GOOS" = "windows" ]; then
        sha256sum "$RELEASE_DIR/$OUTPUT" > "$RELEASE_DIR/$OUTPUT.sha256"
    else
        shasum -a 256 "$RELEASE_DIR/$OUTPUT" > "$RELEASE_DIR/$OUTPUT.sha256"
    fi
done

# Create release notes
cat > "$RELEASE_DIR/RELEASE_NOTES.md" << EOF
# Gitty $VERSION

## Release Information
- **Version**: $VERSION
- **Build Date**: $BUILD_DATE
- **Git Commit**: $GIT_COMMIT

## Downloads
EOF

# Add download links to release notes
for file in "$RELEASE_DIR"/gitty-*; do
    if [[ "$file" != *.sha256 ]]; then
        filename=$(basename "$file")
        echo "- [$filename]($filename)" >> "$RELEASE_DIR/RELEASE_NOTES.md"
    fi
done

echo -e "${GREEN}Release created successfully!${NC}"
echo "Release files are in: $RELEASE_DIR"
echo ""
echo "To create a GitHub release:"
echo "1. Create a new release on GitHub with tag: $VERSION"
echo "2. Upload the files from: $RELEASE_DIR"
echo "3. Use the content from: $RELEASE_DIR/RELEASE_NOTES.md" 