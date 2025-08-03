package main

import (
	"context"
	"fmt"
	"queue-pop-notification/internal/config"
	"queue-pop-notification/internal/service"
	"queue-pop-notification/internal/wow"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading config.yaml")
	}

	callbacks := wow.EventCallbacks{
		OnPvPQueuePop: func(mode wow.PvPMode, details map[string]string) {
			runtime.EventsEmit(a.ctx, "OnPvPQueuePop", map[string]interface{}{
				"mode":      mode,
				"details":   details,
				"timestamp": time.Now().Format("15:04:05"),
			})
		},
	}

	// Start the service with event callbacks
	go func() {

		service.Run(ctx, cfg, &callbacks)
	}()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
