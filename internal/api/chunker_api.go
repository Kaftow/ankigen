package api

import (
	"ankigen/internal/service/chunker"
	"ankigen/internal/service/chunker/types"
	"context"
	"errors"
	"log"
	"time"
)

// ChunkSplitTimeout defines the maximum duration allowed for a single chunking task
const ChunkSplitTimeout = 20 * time.Second

// ChunkJob defines a single chunking task
type ChunkJob struct {
	Text   string            // Text to chunk
	Config types.ChunkConfig // Chunker configuration
	Result chan []types.Chunk
	Err    chan error
	Ctx    context.Context
}

// ChunkerAPI struct
type ChunkerAPI struct {
	ctx          context.Context
	chunkService *chunker.ChunkService
	chunkQueue   chan ChunkJob
}

// NewChunkerAPI creates a new ChunkerAPI application struct
func NewChunkerAPI(poolSize int, queueSize int) *ChunkerAPI {
	cs := chunker.NewChunkService()

	api := &ChunkerAPI{
		chunkService: cs,
		chunkQueue:   make(chan ChunkJob, queueSize),
	}

	// Start the worker pool
	for i := range make([]struct{}, poolSize) {
		go api.chunkWorker(i)
	}

	return api
}

// Startup saves the wails context for UI interactions if needed
func (a *ChunkerAPI) Startup(ctx context.Context) {
	a.ctx = ctx
}

// chunkWorker continuously processes chunking jobs from the queue
func (a *ChunkerAPI) chunkWorker(id int) {
	for job := range a.chunkQueue {
		a.handleChunkJob(id, job)
	}
}

// handleChunkJob executes a single ChunkJob with panic protection
func (a *ChunkerAPI) handleChunkJob(workerID int, job ChunkJob) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Chunk worker %d recovered from panic: %v", workerID, r)
			job.Err <- errors.New("internal error during chunking")
		}
	}()

	log.Printf("Chunk worker %d processing job", workerID)

	chunks, err := a.chunkService.SplitText(job.Text, job.Config)
	if err != nil {
		job.Err <- err
	} else {
		job.Result <- chunks
	}
}

// SplitText splits the provided text using the given chunk config.
// Uses a worker pool to control concurrency while maintaining synchronous return.
func (a *ChunkerAPI) SplitText(text string, cfg types.ChunkConfig) ([]types.Chunk, error) {
	if a.chunkQueue == nil {
		return nil, errors.New("chunker queue not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), ChunkSplitTimeout)
	defer cancel()

	job := ChunkJob{
		Text:   text,
		Config: cfg,
		Result: make(chan []types.Chunk, 1),
		Err:    make(chan error, 1),
		Ctx:    ctx,
	}

	// Submit the job to the queue
	select {
	case a.chunkQueue <- job:
		// Wait for worker result
		select {
		case res := <-job.Result:
			return res, nil
		case err := <-job.Err:
			return nil, err
		case <-ctx.Done():
			return nil, errors.New("chunking timed out")
		}
	default:
		// Queue is full
		return nil, errors.New("chunker queue full")
	}
}
