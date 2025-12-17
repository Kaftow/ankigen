package pandoc

import (
	"os"
	"path/filepath"
	"runtime"
)

func PandocPath() string {
	exePath, err := os.Executable()
	if err != nil {
		return ""
	}
	exeDir := filepath.Dir(exePath)
	return filepath.Join(exeDir, "pandoc", relativePath())
}

func relativePath() string {
	osType := runtime.GOOS
	arch := runtime.GOARCH

	switch osType {
	case "linux":
		if arch == "amd64" {
			return "pandoc-3.8.3-linux-amd64/bin/pandoc"
		} else if arch == "arm64" {
			return "pandoc-3.8.3-linux-arm64/bin/pandoc"
		} else {
			return ""
		}
	case "darwin":
		if arch == "amd64" {
			return "pandoc-3.8.3-x86_64-macOS/bin/pandoc"
		} else if arch == "arm64" {
			return "pandoc-3.8.3-arm64-macOS/bin/pandoc"
		} else {
			return ""
		}
	case "windows":
		return "pandoc-3.8.3-windows-x86_64/pandoc.exe"
	default:
		return ""
	}
}
