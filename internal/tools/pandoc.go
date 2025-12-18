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
		switch arch {
		case "amd64":
			return "pandoc-3.8.3-linux-amd64/pandoc"
		case "arm64":
			return "pandoc-3.8.3-linux-arm64/pandoc"
		default:
			return ""
		}
	case "darwin":
		switch arch {
		case "amd64":
			return "pandoc-3.8.3-x86_64-macOS/pandoc"
		case "arm64":
			return "pandoc-3.8.3-arm64-macOS/pandoc"
		default:
			return ""
		}
	case "windows":
		return "pandoc-3.8.3-windows-x86_64/pandoc.exe"
	default:
		return ""
	}
}
