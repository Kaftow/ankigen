package main

import (
	"context"
	"ankigen/internal/extractor"
)

// App struct
type App struct {
	ctx context.Context
	extractorManager *extractor.ExtractorManager
}

// NewApp creates a new App application struct
func NewApp() *App {
	em := extractor.NewExtractorManager()
	em.RegisterAll()
	return &App{
		extractorManager: em,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// ExtractText extracts text from a file at the given path
func (a *App) ExtractText(path string) (string, error) {
	text, err := a.extractorManager.Extract(path)
	if err != nil {
		return "", err
	}
	return text, nil
}

// GetSupportedExtensions returns all supported file extensions
func (a *App) GetSupportedExtensions() []string {
	return a.extractorManager.SupportedExtensions()
}
