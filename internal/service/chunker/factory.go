package chunker

import (
	"ankigen/internal/service/chunker/impl"
	"ankigen/internal/service/chunker/types"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type TokenParams struct {
	MaxTokens    int    `mapstructure:"maxTokens"`
	EncodingName string `mapstructure:"encodingName"`
}

type FixedLengthParams struct {
	MaxChars int `mapstructure:"maxChars"`
}

func CreateChunker(cfg types.ChunkConfig) (types.Chunker, error) {
	switch cfg.Strategy {
	case "fixedLength":
		var params FixedLengthParams
		if err := mapstructure.Decode(cfg.Params, &params); err != nil {
			return nil, fmt.Errorf("invalid fixedLength params: %w", err)
		}
		return impl.NewFixedLengthChunker(params.MaxChars), nil

	case "token":
		var params TokenParams
		if err := mapstructure.Decode(cfg.Params, &params); err != nil {
			return nil, fmt.Errorf("invalid token params: %w", err)
		}
		return impl.NewTokenChunker(params.MaxTokens, params.EncodingName)

	default:
		return nil, fmt.Errorf("unsupported strategy: %s", cfg.Strategy)
	}
}
