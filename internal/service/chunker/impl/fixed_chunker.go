package impl

type FixedLengthChunker struct {
	MaxChars int
}

func NewFixedLengthChunker(maxChars int) *FixedLengthChunker {
	return &FixedLengthChunker{
		MaxChars: maxChars,
	}
}

func (c *FixedLengthChunker) Split(text string) ([]RawChunk, error) {
	var chunks []RawChunk

	runes := []rune(text) // unicode-safe
	for i := 0; i < len(runes); i += c.MaxChars {
		end := i + c.MaxChars
		if end > len(runes) {
			end = len(runes)
		}

		chunks = append(chunks, RawChunk{
			Text:  string(runes[i:end]),
			Start: i,
			End:   end,
			Meta:  map[string]interface{}{},
		})
	}

	return chunks, nil
}
