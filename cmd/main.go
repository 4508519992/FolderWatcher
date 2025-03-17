package main

import (
	"fmt"
	"folder-watcher/internal/watcher"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please provide a folder path")
	}

	folder := os.Args[1]
	w, err := watcher.NewWatcher(folder)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err := w.Watch()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for event := range w.EventChan {
		fmt.Printf("new folder: %s\n", event)
	}
}
