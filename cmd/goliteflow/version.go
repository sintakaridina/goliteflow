package main

import "fmt"

// Version information
var (
	Version   = "1.0.0"
	BuildTime = "unknown"
	GitCommit = "unknown"
	GoVersion = "unknown"
)

// GetVersion returns the version information
func GetVersion() string {
	return fmt.Sprintf("GoliteFlow v%s (build: %s, commit: %s, go: %s)", 
		Version, BuildTime, GitCommit, GoVersion)
}

// PrintVersion prints the version information
func PrintVersion() {
	fmt.Println(GetVersion())
}
