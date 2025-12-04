package types

type RawChunk struct {
	Text  string
	Start int
	End   int
	Meta  map[string]any
}

type Chunker interface {
	Split(text string) ([]RawChunk, error)
}

type Chunk struct {
	ID    string
	Text  string
	Start int
	End   int
	Meta  map[string]any
}

type ChunkConfig struct {
	Strategy string
	Params   map[string]any
}
