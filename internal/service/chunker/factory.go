package chunker

import (
	"ankigen/internal/service/chunker/impl"
	"fmt"
)

func CreateChunker(cfg ChunkConfig) (Chunker, error) {
	switch cfg.Strategy {

	case "fixedLength":
		maxChars, ok := cfg.Params["maxChars"].(int)
		if !ok {
			return nil, fmt.Errorf("invalid maxChars param")
		}
		return impl.NewFixedLengthChunker(maxChars), nil
	case "token":
		maxTokens, ok := cfg.Params["maxTokens"].(int)
		if !ok {
			return nil, fmt.Errorf("invalid maxToken param")
		}
		encodingName, ok := cfg.Params["encodingName"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid encodingName param")
		}
		return impl.NewTokenChunker(maxTokens, encodingName)

	}

	return nil, fmt.Errorf("unsupported strategy: %s", cfg.Strategy)
}
