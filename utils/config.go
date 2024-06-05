package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	APIKey        string        `mapstructure:"API_KEY"`
	BaseURL       string        `mapstructure:"BASE_URL"`
	DurationToken time.Duration `mapstructure:"DURATION_TOKEN"`
}

func NewConfig(path string) (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
