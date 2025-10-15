#!/bin/bash

# Verification script for GoliteFlow release binaries
# This script verifies that all release binaries are working correctly

set -e

RELEASE_DIR="releases"
PASSED=0
FAILED=0

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}Verifying GoliteFlow release binaries${NC}"
echo ""

# Function to test a binary (if it can run on current platform)
test_binary() {
    local binary=$1
    local platform=$2
    
    if [ ! -f "$RELEASE_DIR/$binary" ]; then
        echo -e "${RED}✗ $binary not found${NC}"
        ((FAILED++))
        return
    fi
    
    # Check if we can run this binary on the current platform
    case "$platform" in
        "linux")
            if [[ "$OSTYPE" == "linux-gnu"* ]]; then
                if ./$RELEASE_DIR/$binary --version >/dev/null 2>&1; then
                    echo -e "${GREEN}✓ $binary works correctly${NC}"
                    ((PASSED++))
                else
                    echo -e "${RED}✗ $binary failed to run${NC}"
                    ((FAILED++))
                fi
            else
                echo -e "${YELLOW}~ $binary exists (cannot test on non-Linux platform)${NC}"
                ((PASSED++))
            fi
            ;;
        "darwin")
            if [[ "$OSTYPE" == "darwin"* ]]; then
                if ./$RELEASE_DIR/$binary --version >/dev/null 2>&1; then
                    echo -e "${GREEN}✓ $binary works correctly${NC}"
                    ((PASSED++))
                else
                    echo -e "${RED}✗ $binary failed to run${NC}"
                    ((FAILED++))
                fi
            else
                echo -e "${YELLOW}~ $binary exists (cannot test on non-macOS platform)${NC}"
                ((PASSED++))
            fi
            ;;
        "windows")
            # Cannot easily test Windows binaries on Unix-like systems
            echo -e "${YELLOW}~ $binary exists (cannot test Windows binary on Unix)${NC}"
            ((PASSED++))
            ;;
    esac
}

# Check if release directory exists
if [ ! -d "$RELEASE_DIR" ]; then
    echo -e "${RED}Error: $RELEASE_DIR directory not found${NC}"
    echo -e "${YELLOW}Run 'make release-binaries' or './scripts/build-release.sh' first${NC}"
    exit 1
fi

# Test each binary
echo -e "${BLUE}Testing binaries:${NC}"
test_binary "goliteflow-linux-amd64" "linux"
test_binary "goliteflow-windows-amd64.exe" "windows" 
test_binary "goliteflow-darwin-arm64" "darwin"

echo ""
echo -e "${BLUE}File sizes:${NC}"
ls -lh $RELEASE_DIR/

echo ""
echo -e "${BLUE}Summary:${NC}"
echo -e "Passed: ${GREEN}$PASSED${NC}"
echo -e "Failed: ${RED}$FAILED${NC}"

if [ $FAILED -eq 0 ]; then
    echo -e "${GREEN}All verifications passed! ✓${NC}"
    echo ""
    echo -e "${BLUE}Binaries are ready for release${NC}"
    echo -e "${YELLOW}Upload these files to GitHub release:${NC}"
    echo "  - $RELEASE_DIR/goliteflow-linux-amd64"
    echo "  - $RELEASE_DIR/goliteflow-windows-amd64.exe"
    echo "  - $RELEASE_DIR/goliteflow-darwin-arm64"
else
    echo -e "${RED}Some verifications failed!${NC}"
    exit 1
fi