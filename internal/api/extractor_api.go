package api

import (
	"ankigen/internal/service/extractor"
	"context"
	"errors"
	"log"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// ExtractTextTimeout defines the maximum duration allowed for a single extraction task
const ExtractTextTimeout = 30 * time.Second

// ExtractJob defines a single extraction task
type ExtractJob struct {
	Path   string          // File path to extract
	Result chan string     // Channel to send the extraction result
	Err    chan error      // Channel to send extraction error
	Ctx    context.Context // Context for cancellation
}

// ExtractorAPI struct
type ExtractorAPI struct {
	ctx              context.Context
	extractorService *extractor.ExtractorService
	extractQueue     chan ExtractJob
}

// NewExtractorAPI creates a new ExtractorAPI application struct
func NewExtractorAPI(poolSize int, queueSize int) *ExtractorAPI {
	em := extractor.NewExtractorService()
	em.RegisterAll()

	api := &ExtractorAPI{
		extractorService: em,
		extractQueue:     make(chan ExtractJob, queueSize),
	}

	// Start the worker pool
	for i := range make([]struct{}, poolSize) {
		go api.extractTextWorker(i)
	}

	return api
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *ExtractorAPI) Startup(ctx context.Context) {
	a.ctx = ctx
}

// extractTextWorker is a single extractTextWorker that continuously processes jobs from the queue
func (a *ExtractorAPI) extractTextWorker(id int) {
	for job := range a.extractQueue {
		a.handleExtractJob(id, job)
	}
}

// handleExtractJob executes a single ExtractJob with panic protection
func (a *ExtractorAPI) handleExtractJob(workerID int, job ExtractJob) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Worker %d recovered from panic: %v", workerID, r)
			job.Err <- errors.New("internal error during extraction")
		}
	}()

	log.Printf("Worker %d processing: %s", workerID, job.Path)

	text, err := a.extractorService.Extract(job.Path)
	if err != nil {
		job.Err <- err
	} else {
		job.Result <- text
	}
}

// ExtractText extracts text from a file at the given path
// Uses a worker pool to control concurrency while maintaining synchronous return
func (a *ExtractorAPI) ExtractText(path string) (string, error) {
	if a.extractQueue == nil {
		return "", errors.New("extractor queue not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), ExtractTextTimeout)
	defer cancel()

	// Create a job for this extraction task
	job := ExtractJob{
		Path:   path,
		Result: make(chan string, 1),
		Err:    make(chan error, 1),
		Ctx:    ctx,
	}

	// Submit the job to the queue
	select {
	case a.extractQueue <- job:
		// Wait for the worker to process and return result
		select {
		case text := <-job.Result:
			return text, nil
		case err := <-job.Err:
			return "", err
		case <-ctx.Done():
			return "", errors.New("extraction timed out")
		}
	default:
		// Queue is full
		return "", errors.New("extract queue full")
	}
}

// GetSupportedExtensions returns all supported file extensions
func (a *ExtractorAPI) GetSupportedExtensions() []string {
	return a.extractorService.SupportedExtensions()
}

func (a *ExtractorAPI) SelectFile() (string, error) {
	if a.ctx == nil {
		log.Printf("wails context not initialized; Startup was not called")
		return "", errors.New("wails context not initialized; Startup was not called")
	}
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select a file",
	})
	if err != nil {
		return "", err
	}
	return file, nil
}
