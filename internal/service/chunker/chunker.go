package chunker

import (
	"ankigen/internal/service/chunker/types"
	"github.com/google/uuid"
	"maps"
)

type ChunkService struct{}

func NewChunkService() *ChunkService {
	return &ChunkService{}
}

// SplitText creates a new chunker based on the request config
func (s *ChunkService) SplitText(text string, cfg types.ChunkConfig) ([]types.Chunk, error) {
	// Create a new chunker for this request
	ch, err := CreateChunker(cfg)
	if err != nil {
		return nil, err
	}

	// Split the text using the newly created chunker
	rawChunks, err := ch.Split(text)
	if err != nil {
		return nil, err
	}

	return RawToChunk(rawChunks, map[string]any{}), nil
}

func RawToChunk(raw []types.RawChunk, newMeta map[string]any) []types.Chunk {
	chunks := make([]types.Chunk, len(raw))
	for i, rc := range raw {
		// Merge existing meta with new meta
		mergedMeta := make(map[string]any)
		maps.Copy(mergedMeta, rc.Meta)
		maps.Copy(mergedMeta, newMeta)
		// Create Chunk with unique ID
		chunks[i] = types.Chunk{
			ID:    uuid.NewString(),
			Text:  rc.Text,
			Start: rc.Start,
			End:   rc.End,
			Meta:  mergedMeta,
		}
	}
	return chunks
}
