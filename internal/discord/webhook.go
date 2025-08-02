package discord

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Webhook struct {
	url         string
	defaultBody string
}

type WebhookOverride struct {
	url  *string
	body *string
}

type WebhookSendOption func() WebhookOverride

func NewWebhook(url string, defaultBody string) *Webhook {
	return &Webhook{url: url, defaultBody: defaultBody}
}

func (w *Webhook) SendNotification(options ...WebhookSendOption) error {
	// Here you would implement the logic to send a notification to the Discord webhook.
	// This is a placeholder implementation.
	fmt.Printf("Sending notification")

	url := w.url
	newBody := w.defaultBody

	for _, opt := range options {
		option := opt()
		if option.url != nil {
			url = *option.url
		}
		if option.body != nil {
			newBody = *option.body
		}
	}

	_, err := http.Post(url, "application/json", strings.NewReader(newBody))
	if err != nil {
		log.Printf("Error sending notification: %v", err)
		return err
	} else {
		log.Println("Notification sent successfully")
	}
	return nil
}
