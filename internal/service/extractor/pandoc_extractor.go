package extractor

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"ankigen/internal/tools"
)

// PandocExtractor handles multiple document types via Pandoc
type PandocExtractor struct {
	PandocPath string
}

// NewPandocExtractor returns a new PandocExtractor.
func NewPandocExtractor() *PandocExtractor {
	pandocPath := tools.PandocPath()
	if pandocPath == "" {
		log.Printf("pandoc not found in PATH")
	}
	return &PandocExtractor{
		PandocPath: pandocPath,
	}
}

// SupportedExtensions returns all supported file extensions.
func (PandocExtractor) SupportedExtensions() []string {
	return []string{
		".txt", ".md",
		".docx",
		".pptx",
		".html", ".htm",
		".epub",
	}
}

// Extract converts supported file types to Markdown using Pandoc.
func (e *PandocExtractor) Extract(path string) (string, error) {
	ext := strings.ToLower(filepath.Ext(path))
	inputFormat := e.extensionToPandocFormat(ext)
	if inputFormat == "" {
		return "", fmt.Errorf("unsupported file extension: %s", ext)
	}

	// Prepare Pandoc command arguments
	args := []string{path, "-f", inputFormat,
		"-t", "gfm", // GitHub-Flavored Markdown
		"--wrap=none", // Disable line wrapping
	}

	// Special handling for PPTX to extract slide content properly
	if ext == ".pptx" {
		args = append(args, "--slide-level=2")
	}

	cmd := exec.Command(e.PandocPath, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		log.Printf("pandoc failed for %s: %v\n%s", path, err, stderr.String())
		return "", fmt.Errorf("pandoc conversion failed: %w", err)
	}

	result := strings.TrimSpace(stdout.String())
	return result, nil
}

// extensionToPandocFormat maps file extensions to Pandoc input formats.
func (e *PandocExtractor) extensionToPandocFormat(ext string) string {
	switch ext {
	case ".txt", ".md":
		return "markdown"
	case ".docx":
		return "docx"
	case ".pptx":
		return "pptx"
	case ".html", ".htm":
		return "html"
	case ".epub":
		return "epub"
	default:
		return ""
	}
}
