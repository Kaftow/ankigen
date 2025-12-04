package impl

import "ankigen/internal/service/chunker/types"

type FixedLengthChunker struct {
	MaxChars int
}

func NewFixedLengthChunker(maxChars int) *FixedLengthChunker {
	return &FixedLengthChunker{
		MaxChars: maxChars,
	}
}

func (c *FixedLengthChunker) Split(text string) ([]types.RawChunk, error) {
	var chunks []types.RawChunk

	runes := []rune(text) // unicode-safe
	for i := 0; i < len(runes); i += c.MaxChars {
		end := i + c.MaxChars
		if end > len(runes) {
			end = len(runes)
		}

		chunks = append(chunks, types.RawChunk{
			Text:  string(runes[i:end]),
			Start: i,
			End:   end,
			Meta:  map[string]any{},
		})
	}

	return chunks, nil
}
