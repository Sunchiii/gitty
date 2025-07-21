#!/bin/bash

# Test script for Gitty installation
set -e

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${YELLOW}Testing Gitty installation scripts...${NC}"

# Test 1: Check if install script exists and is executable
if [ -f "scripts/install.sh" ] && [ -x "scripts/install.sh" ]; then
    echo -e "${GREEN}✅ Install script exists and is executable${NC}"
else
    echo "❌ Install script missing or not executable"
    exit 1
fi

# Test 2: Check if uninstall script exists and is executable
if [ -f "scripts/uninstall.sh" ] && [ -x "scripts/uninstall.sh" ]; then
    echo -e "${GREEN}✅ Uninstall script exists and is executable${NC}"
else
    echo "❌ Uninstall script missing or not executable"
    exit 1
fi

# Test 3: Check if PowerShell script exists
if [ -f "scripts/install.ps1" ]; then
    echo -e "${GREEN}✅ PowerShell install script exists${NC}"
else
    echo "❌ PowerShell install script missing"
    exit 1
fi

# Test 4: Test install script help
if ./scripts/install.sh --help > /dev/null 2>&1; then
    echo -e "${GREEN}✅ Install script help works${NC}"
else
    echo "❌ Install script help failed"
    exit 1
fi

# Test 5: Test version check (should work even without GitHub releases)
if ./scripts/install.sh --version > /dev/null 2>&1; then
    echo -e "${GREEN}✅ Install script version check works${NC}"
else
    echo "❌ Install script version check failed"
    exit 1
fi

# Test 6: Check if main install script exists
if [ -f "install.sh" ] && [ -x "install.sh" ]; then
    echo -e "${GREEN}✅ Main install script exists and is executable${NC}"
else
    echo "❌ Main install script missing or not executable"
    exit 1
fi

echo -e "${GREEN}🎉 All installation script tests passed!${NC}"
echo -e "${YELLOW}Note: These tests verify script existence and basic functionality.${NC}"
echo -e "${YELLOW}Full installation testing requires GitHub releases to be available.${NC}" 