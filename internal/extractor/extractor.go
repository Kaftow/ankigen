package extractor

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"
)

type Extractor interface {
	SupportedExtensions() []string
	Extract(path string) (string, error)
}

// ExtractorManager holds registered extractors and dispatches Extract calls.
type ExtractorManager struct {
	mu         sync.RWMutex
	extractors map[string]Extractor // key: extension (".txt", ".pdf", ...)
}

// NewExtractorManager creates an empty manager.
func NewExtractorManager() *ExtractorManager {
	return &ExtractorManager{
		extractors: make(map[string]Extractor),
	}
}

// Register adds an extractor to the manager for all its supported extensions.
func (m *ExtractorManager) Register(e Extractor) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, ext := range e.SupportedExtensions() {
		m.extractors[ext] = e
	}
}

// RegisterAll registers multiple extractors at once.
func (m *ExtractorManager) RegisterAll() {
	m.mu.Lock()
	defer m.mu.Unlock()
	extractors := []Extractor{
		NewTXTExtractor(),
		NewPDFExtractor(),
		// NewDocxExtractor(),
		// NewHTMLExtractor(),
		// NewMarkdownExtractor()
	}
	for _, e := range extractors {
		for _, ext := range e.SupportedExtensions() {
			m.extractors[ext] = e
		}
	}
}

// SupportedExtensions returns a deduplicated slice of all supported extensions.
func (m *ExtractorManager) SupportedExtensions() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := make([]string, 0, len(m.extractors))
	for ext := range m.extractors {
		result = append(result, ext)
	}
	return result
}

// Extract selects the appropriate extractor by file extension and runs it.
func (m *ExtractorManager) Extract(path string) (string, error) {
	ext := filepath.Ext(path)
	m.mu.RLock()
	e, ok := m.extractors[ext]
	m.mu.RUnlock()
	if !ok {
		log.Printf("unsupported file extension: %s", ext)
		return "", fmt.Errorf("unsupported file extension: %s", ext)
	}
	return e.Extract(path)
}
