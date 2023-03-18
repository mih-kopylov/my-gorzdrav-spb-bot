package config

import (
	"github.com/caarlos0/env/v7"
	"github.com/pkg/errors"
)

type Config struct {
	Token                  string `env:"TELEGRAM_API_TOKEN,required"`
	FirebaseServiceAccount string `env:"FIREBASE_SERVICE_ACCOUNT,required"`
}

func NewConfig() (*Config, error) {
	config, err := readConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
}

func readConfig() (*Config, error) {
	result := &Config{}
	err := env.Parse(result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config")
	}

	return result, nil
}
