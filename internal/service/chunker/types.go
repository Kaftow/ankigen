package chunker

import "ankigen/internal/service/chunker/impl"

type Chunker interface {
	Split(text string) ([]impl.RawChunk, error)
}

type Chunk struct {
	ID   string
	Text string
	Meta map[string]interface{}
}

type ChunkConfig struct {
	Strategy string                 `json:"strategy"`
	Params   map[string]interface{} `json:"params"`
}
