package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PublicApiKey  string `mapstructure:"PUBLIC_API_KEY"`
	PrivateApiKey string `mapstructure:"PRIVATE_API_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)

	viper.SetConfigName("api")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
