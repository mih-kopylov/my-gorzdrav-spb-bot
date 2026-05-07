package config

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"github.com/joomcode/errorx"
)

type Config struct {
	TelegramApiToken string `env:"TELEGRAM_API_TOKEN,required"`
}

func NewConfig() (*Config, error) {
	config, err := readConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
}

func readConfig() (*Config, error) {
	_ = godotenv.Load()

	result := &Config{}
	err := env.Parse(result)
	if err != nil {
		return nil, errorx.EnhanceStackTrace(err, "failed to read config")
	}

	return result, nil
}
