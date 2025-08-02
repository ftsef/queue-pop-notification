package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

// Config represents the structure of the config.yaml file.
type Config struct {
	Discord struct {
		Webhook struct {
			URL  string `yaml:"webhook_url" required:"true"`
			Body string `yaml:"body" required:"true"`
		} `yaml:"webhook"`
	} `yaml:"discord"`
	Wow struct {
		BasePath string `yaml:"base_path" required:"true"`
	} `yaml:"wow"`
}

// LoadConfig reads the configuration from the given path.
func LoadConfig(path string) (*Config, error) {
	config := &Config{}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(c)
}
