package main

import (
	"context"
	"fmt"
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

	callbacks := wow.EventCallbacks{
		OnPvPQueuePop: func(mode wow.PvPMode, details map[string]string) {
			log.Info().Msgf("x")

			runtime.EventsEmit(a.ctx, "OnPvPQueuePop", map[string]interface{}{
				"mode":      mode,
				"details":   details,
				"timestamp": time.Now().Format("15:04:05"),
				"queueType": string(mode),
			})
		},
	}

	service.Run(ctx, "", &callbacks)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
