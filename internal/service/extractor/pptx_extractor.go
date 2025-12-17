package extractor

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"ankigen/internal/tools"
)

// PptxExtractor handles Microsoft PowerPoint (.pptx) files via pandoc
type PptxExtractor struct {
	PandocPath string
}

// NewPptxExtractor returns a new PptxExtractor.
func NewPptxExtractor() *PptxExtractor {
	pandocPath := tools.PandocPath()
	if pandocPath == "" {
		log.Printf("pandoc not found in PATH")
	}
	return &PptxExtractor{
		PandocPath: pandocPath,
	}
}

// SupportedExtensions returns the extensions this extractor supports.
func (PptxExtractor) SupportedExtensions() []string {
	return []string{".pptx"}
}

// Extract converts pptx -> markdown using pandoc and returns the text.
func (e *PptxExtractor) Extract(path string) (string, error) {
	cmd := exec.Command(
		e.PandocPath,
		path,
		"-f", "pptx",
		"-t", "gfm",
		"--slide-level=2",  // extract content from all slide levels
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
