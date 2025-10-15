# Build script for GoliteFlow (Windows PowerShell)
# This script builds binaries for multiple platforms

param(
    [string]$Version = "1.0.0"
)

# Set error action preference
$ErrorActionPreference = "Stop"

# Get build information
$BuildTime = Get-Date -Format "yyyy-MM-dd_HH:mm:ss"
$GitCommit = try { git rev-parse --short HEAD } catch { "unknown" }
$GoVersion = try { (go version).Split(' ')[2] } catch { "unknown" }

Write-Host "Building GoliteFlow v$Version" -ForegroundColor Blue
Write-Host "Build Time: $BuildTime" -ForegroundColor Blue
Write-Host "Git Commit: $GitCommit" -ForegroundColor Blue
Write-Host "Go Version: $GoVersion" -ForegroundColor Blue
Write-Host ""

# Create build directory
$BuildDir = "build"
if (Test-Path $BuildDir) {
    Remove-Item $BuildDir -Recurse -Force
}
New-Item -ItemType Directory -Path $BuildDir | Out-Null

# Build flags
$LdFlags = "-s -w -X main.Version=$Version -X main.BuildTime=$BuildTime -X main.GitCommit=$GitCommit -X main.GoVersion=$GoVersion"

# Platforms to build for
$Platforms = @(
    @{OS="linux"; Arch="amd64"},
    @{OS="linux"; Arch="arm64"},
    @{OS="windows"; Arch="amd64"},
    @{OS="windows"; Arch="arm64"},
    @{OS="darwin"; Arch="amd64"},
    @{OS="darwin"; Arch="arm64"}
)

# Build function
function Build-Platform {
    param($Platform)
    
    $os = $Platform.OS
    $arch = $Platform.Arch
    
    Write-Host "Building for $os/$arch..." -ForegroundColor Yellow
    
    $outputName = "goliteflow"
    if ($os -eq "windows") {
        $outputName = "$outputName.exe"
    }
    
    $outputPath = "$BuildDir\goliteflow_${Version}_${os}_${arch}\$outputName"
    $outputDir = Split-Path $outputPath -Parent
    
    if (!(Test-Path $outputDir)) {
        New-Item -ItemType Directory -Path $outputDir | Out-Null
    }
    
    # Set environment variables
    $env:GOOS = $os
    $env:GOARCH = $arch
    $env:CGO_ENABLED = "0"
    
    # Build
    $buildCmd = "go build -ldflags `"$LdFlags`" -o `"$outputPath`" ./cmd/goliteflow"
    Invoke-Expression $buildCmd
    
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✓ Built $outputPath" -ForegroundColor Green
        
        # Create archive
        $archiveName = "goliteflow_${Version}_${os}_${arch}"
        $archivePath = "$BuildDir\$archiveName"
        
        if ($os -eq "windows") {
            Compress-Archive -Path "$BuildDir\goliteflow_${Version}_${os}_${arch}" -DestinationPath "$archivePath.zip" -Force
        } else {
            # Use tar for non-Windows platforms (requires WSL or Git Bash)
            $tarCmd = "tar -czf `"$archivePath.tar.gz`" -C `"$BuildDir`" `"goliteflow_${Version}_${os}_${arch}`""
            Invoke-Expression $tarCmd
        }
        
        Write-Host "✓ Created archive for $os/$arch" -ForegroundColor Green
    } else {
        Write-Host "✗ Failed to build for $os/$arch" -ForegroundColor Red
        exit 1
    }
}

# Build for all platforms
foreach ($platform in $Platforms) {
    Build-Platform $platform
}

# Create checksums
Write-Host "Creating checksums..." -ForegroundColor Yellow
$checksumFile = "$BuildDir\goliteflow_${Version}_checksums.txt"
Get-ChildItem $BuildDir -Filter "*.tar.gz", "*.zip" | ForEach-Object {
    $hash = Get-FileHash $_.FullName -Algorithm SHA256
    "$($hash.Hash)  $($_.Name)" | Add-Content $checksumFile
}

Write-Host "✓ Created checksums" -ForegroundColor Green

# Create release notes
Write-Host "Creating release notes..." -ForegroundColor Yellow
$releaseNotes = @"
# GoliteFlow v$Version

## Downloads

### Linux
- **AMD64**: [goliteflow_${Version}_linux_amd64.tar.gz](goliteflow_${Version}_linux_amd64.tar.gz)
- **ARM64**: [goliteflow_${Version}_linux_arm64.tar.gz](goliteflow_${Version}_linux_arm64.tar.gz)

### Windows
- **AMD64**: [goliteflow_${Version}_windows_amd64.zip](goliteflow_${Version}_windows_amd64.zip)
- **ARM64**: [goliteflow_${Version}_windows_arm64.zip](goliteflow_${Version}_windows_arm64.zip)

### macOS
- **AMD64**: [goliteflow_${Version}_darwin_amd64.tar.gz](goliteflow_${Version}_darwin_amd64.tar.gz)
- **ARM64**: [goliteflow_${Version}_darwin_arm64.tar.gz](goliteflow_${Version}_darwin_arm64.tar.gz)

## Installation

### Linux/macOS
``````bash
# Download and extract
wget https://github.com/sintakaridina/goliteflow/releases/download/v$Version/goliteflow_${Version}_linux_amd64.tar.gz
tar -xzf goliteflow_${Version}_linux_amd64.tar.gz
sudo mv goliteflow_${Version}_linux_amd64/goliteflow /usr/local/bin/
``````

### Windows
``````powershell
# Download and extract
Invoke-WebRequest -Uri "https://github.com/sintakaridina/goliteflow/releases/download/v$Version/goliteflow_${Version}_windows_amd64.zip" -OutFile "goliteflow.zip"
Expand-Archive -Path "goliteflow.zip" -DestinationPath "."
# Add to PATH or move to desired location
``````

## Verification

``````bash
goliteflow --version
``````

Expected output:
``````
GoliteFlow v$Version (build: $BuildTime, commit: $GitCommit, go: $GoVersion)
``````

## Checksums

``````
$(Get-Content $checksumFile)
``````
"@

$releaseNotes | Out-File -FilePath "$BuildDir\RELEASE_NOTES.md" -Encoding UTF8

Write-Host "✓ Created release notes" -ForegroundColor Green

# Summary
Write-Host ""
Write-Host "Build completed successfully!" -ForegroundColor Green
Write-Host "Build artifacts are in the $BuildDir\ directory" -ForegroundColor Blue
Write-Host ""
Write-Host "Files created:" -ForegroundColor Yellow
Get-ChildItem $BuildDir | Format-Table Name, Length, LastWriteTime
Write-Host ""
Write-Host "To create a release:" -ForegroundColor Blue
Write-Host "1. Create a new tag: git tag v$Version"
Write-Host "2. Push the tag: git push origin v$Version"
Write-Host "3. Upload files from $BuildDir\ to GitHub release"
Write-Host ""
