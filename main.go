package main

import (
	"github.com/richardamare/gitai/cmd"
)

var version string // This variable will be populated by the linker flag

func main() {
	if version == "" {
		version = "dev" // Default for local builds
	}
	cmd.AppVersion = version // Pass the linker-provided version to the cmd package
	cmd.Execute()
}