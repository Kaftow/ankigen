package extractor

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type PDFExtractor struct{}

func NewPDFExtractor() *PDFExtractor {
	return &PDFExtractor{}
}

func (PDFExtractor) SupportedExtensions() []string {
	return []string{".pdf"}
}

func (PDFExtractor) Extract(path string) (string, error) {
	cmd := exec.Command("pdftotext", path, "-")
	output, err := cmd.Output()
	if err != nil {
		log.Printf("error running pdftotext: %v", err)
		return "", fmt.Errorf("error running pdftotext: %w", err)
	}

	return strings.TrimSpace(string(output)), nil
}
