package tools

import (
	"log"
	"os/exec"
	"sync"
)

var (
	pandocPath string
	once       sync.Once
)

func PandocPath() string {
	once.Do(func() {
		path, err := exec.LookPath("pandoc")
		if err != nil {
			log.Printf("pandoc not found in PATH; Extractors will not function")
			return
		}
		pandocPath = path
	})
	return pandocPath
}