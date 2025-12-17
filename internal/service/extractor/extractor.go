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

// ExtractorService holds registered extractors and dispatches Extract calls.
type ExtractorService struct {
	mu         sync.RWMutex
	extractors map[string]Extractor // key: extension (".txt", ".pdf", ...)
}

// NewExtractorService creates an empty service.
func NewExtractorService() *ExtractorService {
	return &ExtractorService{
		extractors: make(map[string]Extractor),
	}
}

// Register adds an extractor to the service for all its supported extensions.
func (m *ExtractorService) Register(e Extractor) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, ext := range e.SupportedExtensions() {
		m.extractors[ext] = e
	}
}

// RegisterAll registers multiple extractors at once.
func (m *ExtractorService) RegisterAll() {
	m.mu.Lock()
	defer m.mu.Unlock()
	extractors := []Extractor{
		NewTXTExtractor(),
		NewPDFExtractor(),
		NewDocxExtractor(),
		NewPptxExtractor(),
	}
	for _, e := range extractors {
		for _, ext := range e.SupportedExtensions() {
			m.extractors[ext] = e
		}
	}
}

// SupportedExtensions returns a deduplicated slice of all supported extensions.
func (m *ExtractorService) SupportedExtensions() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := make([]string, 0, len(m.extractors))
	for ext := range m.extractors {
		result = append(result, ext)
	}
	return result
}

// Extract selects the appropriate extractor by file extension and runs it.
func (m *ExtractorService) Extract(path string) (string, error) {
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
