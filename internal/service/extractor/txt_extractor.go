package extractor

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// TXTExtractor handles plain text files (and simple markdown)
type TXTExtractor struct{}

// NewTXTExtractor returns a new TXTExtractor.
func NewTXTExtractor() *TXTExtractor { return &TXTExtractor{} }

// SupportedExtensions returns the extensions this extractor supports.
func (TXTExtractor) SupportedExtensions() []string {
	return []string{".txt", ".md", ".text"}
}

// Extract reads the file, attempts to convert to UTF-8 and returns the text.
func (TXTExtractor) Extract(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("error reading text file %s: %v", path, err)
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	s, convErr := ToUTF8(data)
	if convErr != nil {
		log.Printf("encoding conversion failed for %s: %v", path, convErr)
		return "", fmt.Errorf("failed to convert file encoding: %w", convErr)
	}

	return strings.TrimSpace(s), nil
}
