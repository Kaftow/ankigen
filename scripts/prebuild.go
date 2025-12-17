package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// This script is a Wails prebuild hook.
// It copies all files from the local "pandoc" directory in the project root
// to "build/bin/pandoc/", so the built application can access Pandoc binaries.
// If the target directory exists, it will be cleared first.

func main() {
	// Get current working directory (assumed to be project root)
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting working directory: %v\n", err)
		os.Exit(1)
	}

	srcDir := filepath.Join(rootDir, "pandoc")            // source pandoc directory
	dstDir := filepath.Join(rootDir, "build", "bin", "pandoc") // target build directory

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
