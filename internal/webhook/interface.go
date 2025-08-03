package webhook

type Webhook interface {
	SendNotification(eventType string, details map[string]string) error
}
