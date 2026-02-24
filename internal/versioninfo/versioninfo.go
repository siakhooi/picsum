/*
Package versioninfo return version
*/
package versioninfo

import "fmt"

var (
	// Version holds the version string of the application.
	Version = "0.0.0"
	// Date holds the build date of the application.
	Date = "unknown"
	// Commit holds the git commit hash of the application.
	Commit = "unknown"
)

// PrintBuildInfo outputs the build information.
func PrintBuildInfo() {
	fmt.Printf("Version: %s\nCommit: %s\nBuildDate: %s\n", Version, Commit, Date)
}
