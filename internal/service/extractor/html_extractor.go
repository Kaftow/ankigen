package extractor

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"ankigen/internal/tools"
)

// HTMLExtractor handles HTML (.html, .htm) files via pandoc
type HTMLExtractor struct {
	PandocPath string
}

// NewHTMLExtractor returns a new HTMLExtractor.
func NewHTMLExtractor() *HTMLExtractor {
	pandocPath := tools.PandocPath()
	if pandocPath == "" {
		log.Printf("pandoc not found in PATH")
	}
	return &HTMLExtractor{
		PandocPath: pandocPath,
	}
}

// SupportedExtensions returns the extensions this extractor supports.
func (HTMLExtractor) SupportedExtensions() []string {
	return []string{".html", ".htm"}
}

// Extract converts HTML -> Markdown using pandoc and returns the text.
func (e *HTMLExtractor) Extract(path string) (string, error) {
	cmd := exec.Command(
		e.PandocPath,
		path,
		"-f", "html",
		"-t", "gfm",       // GitHub Flavored Markdown
		"--wrap=none",     // do not wrap lines
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
