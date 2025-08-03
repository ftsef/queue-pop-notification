package config

import (
	"github.com/go-playground/validator/v10"
)

const DISCORD = "discord"
const NTFY = "ntfy"

type WebhookSender struct {
	Name        string            `yaml:"name" required:"true"`
	URL         string            `yaml:"url" required:"true"`
	Body        string            `yaml:"body" required:"true"`
	Header      map[string]string `yaml:"header,omitempty"`
	RequestType string            `yaml:"request_type" required:"true"` // e.g., "POST", "GET"
	Type        string            `yaml:"type" required:"true"`         // e.g., "discord", "ntfy", etc.
	Enabled     bool              `yaml:"enabled" required:"true"`
}

// Config represents the structure of the config.yaml file.
type Config struct {
	Webhooks []WebhookSender `yaml:"webhooks" required:"true"`
	Wow      struct {
		BasePath string `yaml:"base_path" required:"true"`
	} `yaml:"wow"`
}

func (c *Config) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(c)
}
