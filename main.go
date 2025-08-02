package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"solo-queue-pop/internal/discord"
	"solo-queue-pop/internal/watcher"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Received interrupt signal. Shutting down gracefully...")
		cancel()
	}()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	webhookURL := os.Getenv("WEBHOOK_URL")
	webhookBody := os.Getenv("WEBHOOK_BODY")
	wowBasePath := os.Getenv("WOW_BASE_PATH")

	fmt.Println("Solo-Queue-Pop Notification Service started")
	fmt.Println("Environment variables loaded:")
	fmt.Println("WOW_BASE_PATH:", wowBasePath)

	screenshotDir := filepath.Join(wowBasePath, "_retail_", "Screenshots")
	fmt.Printf("Watching directory: %s\n", screenshotDir)
	fmt.Println("Press Ctrl+C to stop...")

	if _, err := os.Stat(screenshotDir); os.IsNotExist(err) {
		log.Fatalf("Screenshot directory does not exist: %s", screenshotDir)
	}
	if webhookURL == "" {
		log.Fatal("WEBHOOK_URL environment variable is not set")
	}
	if webhookBody == "" {
		log.Println("WEBHOOK_BODY environment variable is not set")
	}

	webhook := discord.NewWebhook(webhookURL, webhookBody)
	sendWebhookCallback := func(filename string) {
		webhook.SendNotification() // or specify option for dynamic body
	}

	w := watcher.NewWatcher(screenshotDir, sendWebhookCallback)
	if err != nil {
		log.Fatal("Error starting directory watcher:", err)
	}

	log.Println(w.Start(ctx))
}
