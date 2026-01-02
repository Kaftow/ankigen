package model

import "gorm.io/gorm"

// The protocol family of an LLM API.
type APIStyle string

const (
	// OpenAI or OpenAI-compatible APIs
	APIStyleOpenAI APIStyle = "openai"

	// Google's Gemini APIs
	APIStyleGemini APIStyle = "gemini"

	// Anthropic Claude native APIs
	APIStyleAnthropic APIStyle = "anthropic"

	// Local or self-hosted Ollama servers
	APIStyleOllama APIStyle = "ollama"

	// Fallback for custom or experimental styles that are not yet mapped.
	APIStyleCustom APIStyle = "custom"
)

// LLMConfig stores provider connection settings.
type LLMConfig struct {
	gorm.Model
	UserID    uint     `gorm:"index"`
	Provider  string   `gorm:"uniqueIndex;size:50"` // Provider name (e.g., OpenAI, Gemini)
	APIKey    string   `gorm:"size:255"`
	BaseURL   string   `gorm:"size:255"`
	ModelName string   `gorm:"size:100"`
	APIStyle  APIStyle `gorm:"size:20"`       // Protocol family (e.g., openai, gemini)
	IsActive  bool     `gorm:"default:false"` // Whether this config is currently active
}
