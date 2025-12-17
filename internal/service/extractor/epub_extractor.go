package extractor

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"ankigen/internal/tools"
)

// EPUBExtractor handles EPUB (.epub) files via pandoc
type EPUBExtractor struct {
	PandocPath string
}

// NewEPUBExtractor returns a new EPUBExtractor.
func NewEPUBExtractor() *EPUBExtractor {
	pandocPath := tools.PandocPath()
	if pandocPath == "" {
		log.Printf("pandoc not found in PATH")
	}
	return &EPUBExtractor{
		PandocPath: pandocPath,
	}
}

// SupportedExtensions returns the extensions this extractor supports.
func (EPUBExtractor) SupportedExtensions() []string {
	return []string{".epub"}
}

// Extract converts EPUB -> Markdown using pandoc and returns the text.
func (e *EPUBExtractor) Extract(path string) (string, error) {
	cmd := exec.Command(
		e.PandocPath,
		path,
		"-f", "epub",
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
