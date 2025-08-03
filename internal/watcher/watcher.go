package watcher

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog/log"
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

	log.Info().Msgf("Started watching for .tga files in: %s", w.dir)

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
				if event.Op == fsnotify.Create {
					log.Info().Msgf("New .tga file detected: %s", event.Name)
					w.callback(event.Name)

					err := os.Remove(event.Name)
					if err != nil {
						log.Warn().Err(err).Msgf("Failed to delete file: %s", event.Name)
					}
				}

			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return err
			}
			log.Error().Err(err).Msg("Watcher error")
		}
	}
}
