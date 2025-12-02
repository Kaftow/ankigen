package api

import (
	"ankigen/internal/service/extractor"
	"context"
	"errors"
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// ExtractorAPI struct
type ExtractorAPI struct {
	ctx              context.Context
	extractorService *extractor.ExtractorService
}

// NewExtractorAPI creates a new ExtractorAPI application struct
func NewExtractorAPI() *ExtractorAPI {
	em := extractor.NewExtractorService()
	em.RegisterAll()
	return &ExtractorAPI{
		extractorService: em,
	}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *ExtractorAPI) Startup(ctx context.Context) {
	a.ctx = ctx
}

// ExtractText extracts text from a file at the given path
func (a *ExtractorAPI) ExtractText(path string) (string, error) {
	text, err := a.extractorService.Extract(path)
	if err != nil {
		return "", err
	}
	return text, nil
}

// GetSupportedExtensions returns all supported file extensions
func (a *ExtractorAPI) GetSupportedExtensions() []string {
	return a.extractorService.SupportedExtensions()
}

func (a *ExtractorAPI) SelectFile() (string, error) {
	if a.ctx == nil {
		log.Printf("wails context not initialized; Startup was not called")
        return "", errors.New("wails context not initialized; Startup was not called")
    }
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select a file",
	})
	if err != nil {
		return "", err
	}
	return file, nil
}
