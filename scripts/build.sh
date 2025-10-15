#!/bin/bash

# Build script for GoliteFlow
# This script builds binaries for multiple platforms

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Version information
VERSION=${1:-"1.0.0"}
BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
GO_VERSION=$(go version | awk '{print $3}')

echo -e "${BLUE}Building GoliteFlow v${VERSION}${NC}"
echo -e "${BLUE}Build Time: ${BUILD_TIME}${NC}"
echo -e "${BLUE}Git Commit: ${GIT_COMMIT}${NC}"
echo -e "${BLUE}Go Version: ${GO_VERSION}${NC}"
echo ""

# Create build directory
BUILD_DIR="build"
mkdir -p ${BUILD_DIR}

# Build flags
LDFLAGS="-s -w -X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -X main.GitCommit=${GIT_COMMIT} -X main.GoVersion=${GO_VERSION}"

# Platforms to build for
PLATFORMS=(
    "linux/amd64"
    "linux/arm64"
    "windows/amd64"
    "windows/arm64"
    "darwin/amd64"
    "darwin/arm64"
)

# Build function
build_platform() {
    local platform=$1
    local os=$(echo $platform | cut -d'/' -f1)
    local arch=$(echo $platform | cut -d'/' -f2)
    
    echo -e "${YELLOW}Building for ${os}/${arch}...${NC}"
    
    local output_name="goliteflow"
    if [ "$os" = "windows" ]; then
        output_name="${output_name}.exe"
    fi
    
    local output_path="${BUILD_DIR}/goliteflow_${VERSION}_${os}_${arch}/${output_name}"
    
    GOOS=$os GOARCH=$arch CGO_ENABLED=0 go build \
        -ldflags "${LDFLAGS}" \
        -o "${output_path}" \
        ./cmd/goliteflow
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ Built ${output_path}${NC}"
        
        # Create archive
        cd ${BUILD_DIR}
        if [ "$os" = "windows" ]; then
            zip -r "goliteflow_${VERSION}_${os}_${arch}.zip" "goliteflow_${VERSION}_${os}_${arch}/"
        else
            tar -czf "goliteflow_${VERSION}_${os}_${arch}.tar.gz" "goliteflow_${VERSION}_${os}_${arch}/"
        fi
        cd ..
        
        echo -e "${GREEN}✓ Created archive for ${os}/${arch}${NC}"
    else
        echo -e "${RED}✗ Failed to build for ${os}/${arch}${NC}"
        exit 1
    fi
}

# Build for all platforms
for platform in "${PLATFORMS[@]}"; do
    build_platform $platform
done

# Create checksums
echo -e "${YELLOW}Creating checksums...${NC}"
cd ${BUILD_DIR}
sha256sum *.tar.gz *.zip > "goliteflow_${VERSION}_checksums.txt"
cd ..

echo -e "${GREEN}✓ Created checksums${NC}"

# Create release notes
echo -e "${YELLOW}Creating release notes...${NC}"
cat > ${BUILD_DIR}/RELEASE_NOTES.md << EOF
# GoliteFlow v${VERSION}

## Downloads

### Linux
- **AMD64**: [goliteflow_${VERSION}_linux_amd64.tar.gz](goliteflow_${VERSION}_linux_amd64.tar.gz)
- **ARM64**: [goliteflow_${VERSION}_linux_arm64.tar.gz](goliteflow_${VERSION}_linux_arm64.tar.gz)

### Windows
- **AMD64**: [goliteflow_${VERSION}_windows_amd64.zip](goliteflow_${VERSION}_windows_amd64.zip)
- **ARM64**: [goliteflow_${VERSION}_windows_arm64.zip](goliteflow_${VERSION}_windows_arm64.zip)

### macOS
- **AMD64**: [goliteflow_${VERSION}_darwin_amd64.tar.gz](goliteflow_${VERSION}_darwin_amd64.tar.gz)
- **ARM64**: [goliteflow_${VERSION}_darwin_arm64.tar.gz](goliteflow_${VERSION}_darwin_arm64.tar.gz)

## Installation

### Linux/macOS
\`\`\`bash
# Download and extract
wget https://github.com/sintakaridina/goliteflow/releases/download/v${VERSION}/goliteflow_${VERSION}_linux_amd64.tar.gz
tar -xzf goliteflow_${VERSION}_linux_amd64.tar.gz
sudo mv goliteflow_${VERSION}_linux_amd64/goliteflow /usr/local/bin/
\`\`\`

### Windows
\`\`\`powershell
# Download and extract
Invoke-WebRequest -Uri "https://github.com/sintakaridina/goliteflow/releases/download/v${VERSION}/goliteflow_${VERSION}_windows_amd64.zip" -OutFile "goliteflow.zip"
Expand-Archive -Path "goliteflow.zip" -DestinationPath "."
# Add to PATH or move to desired location
\`\`\`

## Verification

\`\`\`bash
goliteflow --version
\`\`\`

Expected output:
\`\`\`
GoliteFlow v${VERSION} (build: ${BUILD_TIME}, commit: ${GIT_COMMIT}, go: ${GO_VERSION})
\`\`\`

## Checksums

\`\`\`
$(cat ${BUILD_DIR}/goliteflow_${VERSION}_checksums.txt)
\`\`\`
EOF

echo -e "${GREEN}✓ Created release notes${NC}"

# Summary
echo ""
echo -e "${GREEN}Build completed successfully!${NC}"
echo -e "${BLUE}Build artifacts are in the ${BUILD_DIR}/ directory${NC}"
echo ""
echo -e "${YELLOW}Files created:${NC}"
ls -la ${BUILD_DIR}/
echo ""
echo -e "${BLUE}To create a release:${NC}"
echo -e "1. Create a new tag: \`git tag v${VERSION}\`"
echo -e "2. Push the tag: \`git push origin v${VERSION}\`"
echo -e "3. Upload files from ${BUILD_DIR}/ to GitHub release"
echo ""
