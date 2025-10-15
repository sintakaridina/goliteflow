# Building and Releasing GoliteFlow

This document explains how to build and release GoliteFlow binaries for the platforms mentioned in the README.

## Prerequisites

- Go 1.19 or later installed
- Git (for version information)
- Make (optional, for using Makefile targets)

## Building Release Binaries

### Method 1: Using the Makefile (Recommended)

```bash
make release-binaries
```

This will create binaries in the `releases/` directory with the exact names expected by the README:

- `goliteflow-linux-amd64`
- `goliteflow-windows-amd64.exe`
- `goliteflow-darwin-arm64`

### Method 2: Using the Build Scripts

#### On Linux/macOS:

```bash
chmod +x scripts/build-release.sh
./scripts/build-release.sh
```

#### On Windows:

```powershell
PowerShell -ExecutionPolicy Bypass -File scripts/build-release.ps1
```

### Method 3: Manual Build Commands

```bash
# Create releases directory
mkdir -p releases

# Linux AMD64
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o releases/goliteflow-linux-amd64 ./cmd/goliteflow

# Windows AMD64
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o releases/goliteflow-windows-amd64.exe ./cmd/goliteflow

# macOS ARM64 (Apple Silicon)
GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "-s -w" -o releases/goliteflow-darwin-arm64 ./cmd/goliteflow

# Make Unix binaries executable
chmod +x releases/goliteflow-linux-amd64 releases/goliteflow-darwin-arm64
```

## Manual Release Process

Since we're not using automated GitHub Actions, you'll need to create releases manually:

### Step 1: Build the Binaries

```bash
# Using PowerShell on Windows
PowerShell -ExecutionPolicy Bypass -File scripts/build-release.ps1

# Or using Makefile
make release-binaries
```

### Step 2: Create GitHub Release

1. **Go to GitHub repository** → Releases → "Create a new release"
2. **Choose a tag**: Create new tag (e.g., `v1.0.1`)
3. **Release title**: `GoliteFlow v1.0.1`
4. **Upload binaries**: Drag and drop the 3 files from `releases/` folder:
   - `goliteflow-linux-amd64`
   - `goliteflow-windows-amd64.exe`
   - `goliteflow-darwin-arm64`
5. **Add release notes** (see template below)

### Step 3: Release Notes Template

```markdown
## GoliteFlow v1.0.1

### Download Instructions

#### Linux (AMD64)

\`\`\`bash
curl -L https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-linux-amd64 -o goliteflow
chmod +x goliteflow
\`\`\`

#### Windows (AMD64)

\`\`\`bash
curl -L https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-windows-amd64.exe -o goliteflow.exe
\`\`\`

#### macOS (Apple Silicon)

\`\`\`bash
curl -L https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-darwin-arm64 -o goliteflow
chmod +x goliteflow
\`\`\`
```

## Binary Names and URLs

The binaries are built with specific names to match the download URLs in the README:

| Platform      | Binary Name                    | README URL                                                                                          |
| ------------- | ------------------------------ | --------------------------------------------------------------------------------------------------- |
| Linux AMD64   | `goliteflow-linux-amd64`       | `https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-linux-amd64`       |
| Windows AMD64 | `goliteflow-windows-amd64.exe` | `https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-windows-amd64.exe` |
| macOS ARM64   | `goliteflow-darwin-arm64`      | `https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-darwin-arm64`      |

## Build Flags

The binaries are built with optimization flags:

- `-s`: Strip symbol table and debug information
- `-w`: Strip DWARF debug information
- `CGO_ENABLED=0`: Disable CGO for static linking

This results in smaller, standalone binaries that don't require external dependencies.

## Testing the Binaries

After building, you can test the binaries:

```bash
# Test Linux binary (on Linux)
./releases/goliteflow-linux-amd64 --version

# Test Windows binary (on Windows)
releases\goliteflow-windows-amd64.exe --version

# Test macOS binary (on macOS with Apple Silicon)
./releases/goliteflow-darwin-arm64 --version
```

## File Sizes

Typical binary sizes:

- Linux AMD64: ~8-12 MB
- Windows AMD64: ~8-12 MB
- macOS ARM64: ~8-12 MB

## Troubleshooting

### Go not installed

If you get "go: command not found", install Go from https://golang.org/dl/

### Permission denied (Linux/macOS)

Make sure the scripts are executable:

```bash
chmod +x scripts/build-release.sh
```

### CGO errors

Make sure `CGO_ENABLED=0` is set to build static binaries.

### Cross-compilation issues

Go supports cross-compilation out of the box. If you encounter issues, ensure you're using Go 1.19 or later.

## Contributing

When adding new platforms or modifying the build process:

1. Update this documentation
2. Update the GitHub Actions workflow
3. Update the Makefile targets
4. Update the build scripts
5. Test the build process on multiple platforms
