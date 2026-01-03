package llm

import (
	"ankigen/internal/model"
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/llms/openai"
)

// Factory responsible for creating LLM clients based on database config
type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

func validateLLMConfig(cfg model.LLMConfig) error {
	if cfg.APIStyle == "" {
		return fmt.Errorf("APIStyle is required")
	}
	if cfg.ModelName == "" {
		return fmt.Errorf("model name is required")
	}
	if cfg.BaseURL == "" {
		return fmt.Errorf("base URL is required for Ollama")
	}
	switch cfg.APIStyle {
	case model.APIStyleOpenAI, model.APIStyleGemini, model.APIStyleAnthropic:
		if cfg.APIKey == "" {
			return fmt.Errorf("API key is required for %s", cfg.APIStyle)
		}
	}
	return nil
}

// NewClient takes an LLMConfig and returns a generic llms.Model interface
func (f *Factory) NewClient(cfg model.LLMConfig) (llms.Model, error) {
	ctx := context.Background()

	switch cfg.APIStyle {
	case model.APIStyleOpenAI:
		opts := []openai.Option{
			openai.WithToken(cfg.APIKey),
			openai.WithModel(cfg.ModelName),
		}
		if cfg.BaseURL != "" {
			opts = append(opts, openai.WithBaseURL(cfg.BaseURL))
		}
		return openai.New(opts...)
	case model.APIStyleGemini:
		return googleai.New(ctx,
			googleai.WithAPIKey(cfg.APIKey),
			googleai.WithDefaultModel(cfg.ModelName),
		)

	case model.APIStyleAnthropic:
		opts := []anthropic.Option{
			anthropic.WithToken(cfg.APIKey),
			anthropic.WithModel(cfg.ModelName),
		}
		if cfg.BaseURL != "" {
			opts = append(opts, anthropic.WithBaseURL(cfg.BaseURL))
		}
		return anthropic.New(opts...)

	case model.APIStyleOllama:
		opts := []openai.Option{
			openai.WithBaseURL(cfg.BaseURL),
			openai.WithToken(cfg.APIKey),
			openai.WithModel(cfg.ModelName),
		}
		return openai.New(opts...)

	default:
		return nil, fmt.Errorf("unsupported API style: %s", cfg.APIStyle)
	}
}
