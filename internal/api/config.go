package api

import "github.com/kelseyhightower/envconfig"

// Config holds the application-level configurations.
type Config struct {
	AppPort     string `default:"8080" envconfig:"APP_PORT"`
	ServiceName string `default:"bookmark_service" envconfig:"SERVICE_NAME"`
	InstanceID  string `envconfig:"INSTANCE_ID"`
}

// NewConfig loads configuration from environment variables.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
