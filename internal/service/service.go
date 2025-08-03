package service

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"queue-pop-notification/internal/config"
	"queue-pop-notification/internal/discord"
	"queue-pop-notification/internal/watcher"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func Run(ctx context.Context, cfgFile string) {
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

	cfg, err := config.LoadConfig(cfgFile)
	if err != nil {
		log.Fatal().Msgf("Error loading config.yaml: %v", err)
	}

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

	webhook := discord.NewWebhook(cfg.Discord.Webhook.URL, cfg.Discord.Webhook.Body)
	sendWebhookCallback := func(filename string) {
		webhook.SendNotification() // or specify option for dynamic body
	}

	w := watcher.NewWatcher(screenshotDir, sendWebhookCallback)
	if err != nil {
		log.Fatal().Msgf("Error starting directory watcher: %v", err)
	}

	log.Info().Msgf("Directory watcher started: %v", w.Start(ctx))
}
