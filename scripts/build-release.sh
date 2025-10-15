#!/bin/bash

# Build script for GoliteFlow release binaries
# This script creates the exact binaries referenced in the README

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}Building GoliteFlow release binaries${NC}"
echo ""

# Create releases directory
RELEASE_DIR="releases"
mkdir -p ${RELEASE_DIR}

# Build flags for smaller binaries
LDFLAGS="-s -w"

# Build for Linux AMD64
echo -e "${YELLOW}Building goliteflow-linux-amd64...${NC}"
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build \
    -ldflags "${LDFLAGS}" \
    -o "${RELEASE_DIR}/goliteflow-linux-amd64" \
    ./cmd/goliteflow

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Built goliteflow-linux-amd64${NC}"
else
    echo -e "${RED}✗ Failed to build for Linux AMD64${NC}"
    exit 1
fi

# Build for Windows AMD64
echo -e "${YELLOW}Building goliteflow-windows-amd64.exe...${NC}"
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build \
    -ldflags "${LDFLAGS}" \
    -o "${RELEASE_DIR}/goliteflow-windows-amd64.exe" \
    ./cmd/goliteflow

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Built goliteflow-windows-amd64.exe${NC}"
else
    echo -e "${RED}✗ Failed to build for Windows AMD64${NC}"
    exit 1
fi

# Build for macOS ARM64 (Apple Silicon)
echo -e "${YELLOW}Building goliteflow-darwin-arm64...${NC}"
GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build \
    -ldflags "${LDFLAGS}" \
    -o "${RELEASE_DIR}/goliteflow-darwin-arm64" \
    ./cmd/goliteflow

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Built goliteflow-darwin-arm64${NC}"
else
    echo -e "${RED}✗ Failed to build for macOS ARM64${NC}"
    exit 1
fi

# Make Unix binaries executable
chmod +x "${RELEASE_DIR}/goliteflow-linux-amd64"
chmod +x "${RELEASE_DIR}/goliteflow-darwin-arm64"

echo ""
echo -e "${GREEN}All release binaries built successfully!${NC}"
echo ""
echo -e "${BLUE}Release binaries:${NC}"
ls -la ${RELEASE_DIR}/
echo ""
echo -e "${BLUE}File sizes:${NC}"
du -h ${RELEASE_DIR}/*
echo ""
echo -e "${YELLOW}These binaries match the download URLs in your README:${NC}"
echo -e "- goliteflow-linux-amd64"
echo -e "- goliteflow-windows-amd64.exe"
echo -e "- goliteflow-darwin-arm64"
echo ""