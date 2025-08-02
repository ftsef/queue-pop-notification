package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"solo-queue-pop/internal/config"
	"solo-queue-pop/internal/discord"
	"solo-queue-pop/internal/watcher"
	"syscall"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "qpn",
		Short: "qpn - Queue Pop Notifier - Notifies when a queue pop is detected",
		Long:  `Sends a Discord notification when a queue pop is detected in World of Warcraft.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			run()
		},
	}
)

func main() {

	cobra.MousetrapHelpText = ""
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "Path to the configuration file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Received interrupt signal. Shutting down gracefully...")
		cancel()
	}()

	cfg, err := config.LoadConfig(cfgFile)
	if err != nil {
		log.Fatalf("Error loading config.yaml: %v", err)
	}

	fmt.Println("Solo-Queue-Pop Notification Service started")
	fmt.Println("Configuration loaded from config.yaml")
	fmt.Println("WOW_BASE_PATH:", cfg.Wow.BasePath)

	screenshotDir := filepath.Join(cfg.Wow.BasePath, "_retail_", "Screenshots")
	fmt.Printf("Watching directory: %s\n", screenshotDir)
	fmt.Println("Press Ctrl+C to stop...")

	if _, err := os.Stat(screenshotDir); os.IsNotExist(err) {
		log.Fatalf("Screenshot directory does not exist: %s", screenshotDir)
	}

	webhook := discord.NewWebhook(cfg.Discord.Webhook.URL, cfg.Discord.Webhook.Body)
	sendWebhookCallback := func(filename string) {
		webhook.SendNotification() // or specify option for dynamic body
	}

	w := watcher.NewWatcher(screenshotDir, sendWebhookCallback)
	if err != nil {
		log.Fatal("Error starting directory watcher:", err)
	}

	log.Println(w.Start(ctx))
}
