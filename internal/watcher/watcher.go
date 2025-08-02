package watcher

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	dir      string
	callback func(filename string)
}

func NewWatcher(dir string, callback func(filename string)) *Watcher {
	return &Watcher{dir: dir, callback: callback}
}

func (w *Watcher) Start(ctx context.Context) error {
	// Check if directory exists
	if _, err := os.Stat(w.dir); os.IsNotExist(err) {
		return fmt.Errorf("directory does not exist: %s", w.dir)
	}

	// Create new watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	// Add directory to watcher
	err = watcher.Add(w.dir)
	if err != nil {
		return err
	}

	fmt.Printf("Started watching for .tga files in: %s\n", w.dir)

	// Start listening for events
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case event, ok := <-watcher.Events:
			if !ok {
				return fmt.Errorf("watcher events channel closed")
			}

			// Check if it's a .tga file and a create/write event
			if strings.HasSuffix(strings.ToLower(event.Name), ".tga") {
				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Printf("New .tga file detected: %s\n", event.Name)
					w.callback(event.Name)
				} else if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Printf("Modified .tga file: %s\n", event.Name)
				}
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return err
			}
			log.Printf("Watcher error: %v", err)
		}
	}
}
