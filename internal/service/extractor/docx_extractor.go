package extractor

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"ankigen/internal/tools"
)


// DocxExtractor handles Microsoft Word (.docx) files via pandoc
type DocxExtractor struct {
	PandocPath string
}

// NewDocxExtractor returns a new DocxExtractor.
func NewDocxExtractor() *DocxExtractor {
	pandocPath := tools.PandocPath()
	if pandocPath == "" {
		log.Printf("pandoc not found in PATH")
	}
	return &DocxExtractor{
		PandocPath: pandocPath,
	}
}

// SupportedExtensions returns the extensions this extractor supports.
func (DocxExtractor) SupportedExtensions() []string {
	return []string{".docx"}
}

// Extract converts docx -> markdown using pandoc and returns the text.
func (e *DocxExtractor) Extract(path string) (string, error) {
	cmd := exec.Command(
		e.PandocPath,
		path,
		"-f", "docx",
		"-t", "gfm",
		"--wrap=none",
	)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		log.Printf("pandoc failed for %s: %v\n%s", path, err, stderr.String())
		return "", fmt.Errorf("pandoc conversion failed: %w", err)
	}

	result := strings.TrimSpace(stdout.String())
	return result, nil
}
