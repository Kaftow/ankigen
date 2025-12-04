package impl

import (
	"ankigen/internal/service/chunker/types"
	"github.com/pkoukk/tiktoken-go"
	"github.com/sbt-zyq/tiktoken-go-loader" // offline loader
)

type TokenChunker struct {
	MaxTokens int
	encoding  *tiktoken.Tiktoken
}

func NewTokenChunker(maxTokens int, encodingName string) (*TokenChunker, error) {
	// Use offline BPE loader instead of downloading dictionaries at runtime
	tiktoken.SetBpeLoader(tiktoken_loader.NewOfflineLoader())

	// Load encoding by name (e.g. "gpt-4o", "cl100k_base", etc.)
	enc, err := tiktoken.GetEncoding(encodingName)
	if err != nil {
		return nil, err
	}

	return &TokenChunker{
		MaxTokens: maxTokens,
		encoding:  enc,
	}, nil
}

func (c *TokenChunker) Split(text string) ([]types.RawChunk, error) {
	// Convert input text into token IDs
	tokens := c.encoding.Encode(text, nil, nil)
	var chunks []types.RawChunk
	currentStart := 0

	// Slice tokens into chunks
	for i := 0; i < len(tokens); i += c.MaxTokens {
		end := i + c.MaxTokens
		if end > len(tokens) {
			end = len(tokens)
		}

		// Decode tokens back into text
		chunkText := c.encoding.Decode(tokens[i:end])
		chunkLength := len([]rune(chunkText))

		// Create RawChunk
		chunks = append(chunks, types.RawChunk{
			Text:  chunkText,
			Start: currentStart,
			End:   currentStart + chunkLength,
			Meta: map[string]any{
				"start_token": i,
				"end_token":   end,
			},
		})

		currentStart += chunkLength
	}

	return chunks, nil
}
