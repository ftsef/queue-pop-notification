package webhook

import (
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

type DiscordWebhook struct {
	url         string
	defaultBody string
}

func NewDiscordWebhook(url string, defaultBody string) *DiscordWebhook {
	return &DiscordWebhook{url: url, defaultBody: defaultBody}
}

func (w *DiscordWebhook) SendNotification(eventType string, details map[string]string) error {
	// Here you would implement the logic to send a notification to the Discord webhook.
	// This is a placeholder implementation.

	url := w.url
	newBody := w.defaultBody

	_, err := http.Post(url, "application/json", strings.NewReader(newBody))
	if err != nil {
		log.Error().Err(err).Msg("Error sending notification")
		return err
	}
	return nil
}
