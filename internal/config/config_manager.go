package config

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"

	gap "github.com/muesli/go-app-paths"
	"github.com/rs/zerolog/log"
)

const APPNAME = "queue-pop-notification"

type ConfigManager struct {
	config Config
	path   string
}

func NewConfigManager(path string) *ConfigManager {
	return &ConfigManager{
		config: Config{},
		path:   path,
	}
}

func (cm *ConfigManager) Load() (Config, error) {
	scope := gap.NewScope(gap.User, APPNAME)

	cfgPath, err := scope.ConfigPath("config.yaml")
	if err != nil {
		return Config{}, err
	}

	if cm.path == "" {
		log.Info().Msgf("Config file not set. Using default config %s", cm.path)
		cm.path = cfgPath
	}

	d := path.Dir(cm.path)
	if _, err := os.Stat(d); os.IsNotExist(err) {
		log.Info().Msgf("Creating config directory: %s", d)
		if err := os.MkdirAll(d, 0755); err != nil {
			log.Fatal().Err(err).Msg("Failed to create config directory")
		}
	}

	if _, err := os.Stat(cm.path); os.IsNotExist(err) {
		log.Info().Msgf("Config file not found at %s. Creating new config.", cm.path)
		if err := cm.Save(); err != nil {
			log.Fatal().Err(err).Msg("Failed to create default config file")
		}
	}

	log.Info().Msgf("Loading config from %s", cfgPath)

	return cm.loadConfig(cfgPath)
}

func (cm *ConfigManager) loadConfig(path string) (Config, error) {
	config := Config{}

	file, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (cm *ConfigManager) Save() error {
	data, err := yaml.Marshal(cm.config)
	if err != nil {
		return err
	}

	err = os.WriteFile(cm.path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
