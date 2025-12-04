package impl

type RawChunk struct {
	Text  string
	Start int
	End   int
	Meta  map[string]interface{}
}
