package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

// This script is a Wails prebuild hook.
// It copies all files from the local "pandoc" directory in the project root
// to "build/bin/pandoc/", so the built application can access Pandoc binaries.
// If the target directory exists, it will be cleared first.

func getBinaryFolder() string {
	osType := runtime.GOOS
	arch := runtime.GOARCH

	switch osType {
	case "linux":
		switch arch {
		case "amd64":
			return "pandoc-3.8.3-linux-amd64"
		case "arm64":
			return "pandoc-3.8.3-linux-arm64"
		}
	case "darwin":
		switch arch {
		case "amd64":
			return "pandoc-3.8.3-x86_64-macOS"
		case "arm64":
			return "pandoc-3.8.3-arm64-macOS"
		}
	case "windows":
		return "pandoc-3.8.3-windows-x86_64"
	}
	return ""
}

func main() {
	// Get current working directory (assumed to be project root)
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting working directory: %v\n", err)
		os.Exit(1)
	}

	binaryFolder := getBinaryFolder()

	if binaryFolder == "" {
		fmt.Fprintf(os.Stderr, "Unsupported OS/Architecture: %s/%s\n", runtime.GOOS, runtime.GOARCH)
		os.Exit(1)
	}

	srcDir := filepath.Join(rootDir, "pandoc", binaryFolder)                 // source pandoc binary directory
	dstDir := filepath.Join(rootDir, "build", "bin", "pandoc", binaryFolder) // target build directory

	// Remove target directory if it exists
	os.RemoveAll(dstDir)

	// Create target directory
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating build pandoc directory: %v\n", err)
		os.Exit(1)
	}

	// Walk through the source directory and copy files
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dstDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return err
		}

		mode := info.Mode()
		// Make executable if it's a Pandoc binary
		if info.Name() == "pandoc" || info.Name() == "pandoc.exe" {
			mode |= 0111
		}
		return os.Chmod(dstPath, mode)
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error copying pandoc files: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Pandoc binaries copied to build/bin/pandoc successfully.")
}
