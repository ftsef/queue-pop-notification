package service

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"queue-pop-notification/internal/config"
	"queue-pop-notification/internal/watcher"
	"queue-pop-notification/internal/webhook"
	"queue-pop-notification/internal/wow"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

// EventCallbacks defines callback functions for different events

// GetCurrentTime returns the current time formatted for display
func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Run(ctx context.Context, cfg config.Config, otherCallbacks *wow.EventCallbacks) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Warn().Msg("Received interrupt signal. Shutting down gracefully...")
		cancel()
	}()

	if err := cfg.Validate(); err != nil {
		log.Fatal().Msgf("Configuration validation failed: %v", err)
	}

	log.Info().Msg("Queue-Pop-Notification Service started")
	log.Info().Msg("Configuration loaded from config.yaml")
	log.Info().Msgf("WoW base path: %s", cfg.Wow.BasePath)

	screenshotDir := filepath.Join(cfg.Wow.BasePath, "_retail_", "Screenshots")
	log.Info().Msgf("Watching directory: %s", screenshotDir)
	log.Info().Msg("Press Ctrl+C to stop...")

	if _, err := os.Stat(screenshotDir); os.IsNotExist(err) {
		log.Fatal().Msgf("Screenshot directory does not exist: %s", screenshotDir)
	}

	webhooks := make(map[string]webhook.Webhook)

	for _, hook := range cfg.Webhooks {
		log.Info().Msgf("Configuring webhook: %s", hook.Name)
		switch hook.Type {
		case config.DISCORD:
			webhooks[hook.Name] = webhook.NewDiscordWebhook(hook.URL, hook.Body)
		case config.NTFY:
			log.Warn().Msgf("NTFY webhook type is not implemented yet: %s", hook.Name)
		default:
			log.Warn().Msgf("Unknown webhook type: %s", hook.Name)
		}

	}

	callbacks := wow.EventCallbacks{
		OnPvPQueuePop: func(mode wow.PvPMode, details map[string]string) {
			for whName, wh := range webhooks {
				if wh == nil {
					log.Warn().Msgf("Webhook %s is nil, skipping notification", whName)
					continue
				}

				log.Info().Msgf("Queue pop detected, Mode %s", wh)
				if err := wh.SendNotification(string(mode), details); err != nil {
					log.Error().Err(err).Msgf("Failed to send notification via %s", whName)
				} else {
					log.Info().Msgf("Notification sent successfully via %s", whName)
				}
			}

			if otherCallbacks != nil {
				if otherCallbacks.OnPvPQueuePop != nil {
					otherCallbacks.OnPvPQueuePop(mode, details)
				}
			}
		},
	}

	w := watcher.NewWatcher(screenshotDir, callbacks)
	log.Info().Msgf("Directory watcher started: %v", w.Start(ctx))
}
