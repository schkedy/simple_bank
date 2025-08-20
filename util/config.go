package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DB_DRIVER           string        `mapstructure:"DB_DRIVER"`
	DB_SOURCE           string        `mapstructure:"DB_SOURCE"`
	SERVER_ADDRESS      string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

// LoadConfig считывает настройки из файла или переменных окружения.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
