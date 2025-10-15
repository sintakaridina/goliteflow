# Build script for GoliteFlow release binaries (Windows PowerShell)
# This script creates the exact binaries referenced in the README

param()

# Set error action preference
$ErrorActionPreference = "Stop"

Write-Host "Building GoliteFlow release binaries" -ForegroundColor Blue
Write-Host ""

# Create releases directory
$ReleaseDir = "releases"
if (Test-Path $ReleaseDir) {
    Remove-Item $ReleaseDir -Recurse -Force
}
New-Item -ItemType Directory -Path $ReleaseDir | Out-Null

# Build flags for smaller binaries
$LdFlags = "-s -w"

# Build for Linux AMD64
Write-Host "Building goliteflow-linux-amd64..." -ForegroundColor Yellow
$env:GOOS = "linux"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"
go build -ldflags $LdFlags -o "$ReleaseDir/goliteflow-linux-amd64" ./cmd/goliteflow

if ($LASTEXITCODE -eq 0) {
    Write-Host "✓ Built goliteflow-linux-amd64" -ForegroundColor Green
} else {
    Write-Host "✗ Failed to build for Linux AMD64" -ForegroundColor Red
    exit 1
}

# Build for Windows AMD64
Write-Host "Building goliteflow-windows-amd64.exe..." -ForegroundColor Yellow
$env:GOOS = "windows"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"
go build -ldflags $LdFlags -o "$ReleaseDir/goliteflow-windows-amd64.exe" ./cmd/goliteflow

if ($LASTEXITCODE -eq 0) {
    Write-Host "✓ Built goliteflow-windows-amd64.exe" -ForegroundColor Green
} else {
    Write-Host "✗ Failed to build for Windows AMD64" -ForegroundColor Red
    exit 1
}

# Build for macOS ARM64 (Apple Silicon)
Write-Host "Building goliteflow-darwin-arm64..." -ForegroundColor Yellow
$env:GOOS = "darwin"
$env:GOARCH = "arm64"
$env:CGO_ENABLED = "0"
go build -ldflags $LdFlags -o "$ReleaseDir/goliteflow-darwin-arm64" ./cmd/goliteflow

if ($LASTEXITCODE -eq 0) {
    Write-Host "✓ Built goliteflow-darwin-arm64" -ForegroundColor Green
} else {
    Write-Host "✗ Failed to build for macOS ARM64" -ForegroundColor Red
    exit 1
}

# Clean up environment variables
Remove-Item Env:GOOS
Remove-Item Env:GOARCH
Remove-Item Env:CGO_ENABLED

Write-Host ""
Write-Host "All release binaries built successfully!" -ForegroundColor Green
Write-Host ""
Write-Host "Release binaries:" -ForegroundColor Blue
Get-ChildItem $ReleaseDir
Write-Host ""
Write-Host "File sizes:" -ForegroundColor Blue
Get-ChildItem $ReleaseDir | ForEach-Object { 
    $size = [math]::Round($_.Length / 1MB, 2)
    Write-Host "$($_.Name): $size MB"
}
Write-Host ""
Write-Host "These binaries match the download URLs in your README:" -ForegroundColor Yellow
Write-Host "- goliteflow-linux-amd64"
Write-Host "- goliteflow-windows-amd64.exe"
Write-Host "- goliteflow-darwin-arm64"
Write-Host ""