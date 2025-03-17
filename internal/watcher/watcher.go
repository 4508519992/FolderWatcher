package watcher

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Watcher struct {
	Folder    string
	EventChan chan string
}

func NewWatcher(folderPath string) (*Watcher, error) {
	if info, err := os.Stat(folderPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, errors.New("folder doesn't exist")
		}
		return nil, fmt.Errorf("failed to access folder: %w", err)
	} else if !info.IsDir() {
		return nil, errors.New("path is not a directory")
	}

	return &Watcher{
		Folder:    folderPath,
		EventChan: make(chan string, 100),
	}, nil
}

func (w *Watcher) Watch() error {
	seen := make(map[string]bool)

	for {
		files, err := os.ReadDir(w.Folder)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", w.Folder, err)
			continue
		}

		currentFiles := make(map[string]bool)

		for _, file := range files {
			if file.IsDir() {
				fullPath := filepath.Join(w.Folder, file.Name())
				if !seen[file.Name()] {
					w.EventChan <- fullPath
				}
				currentFiles[file.Name()] = true
			}
		}

		for fileName := range seen {
			if !currentFiles[fileName] {
				delete(seen, fileName)
			}
		}

		seen = currentFiles
		time.Sleep(1 * time.Second)
	}
}
