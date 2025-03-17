package main

import (
	"fmt"
	"folder-watcher/internal/watcher"
	"log"
)

func main() {
	w, err := watcher.NewWatcher("./test")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		w.Watch()
	}()

	for event := range w.EventChan {
		fmt.Printf("new folder: %s\n", event)
	}
}
