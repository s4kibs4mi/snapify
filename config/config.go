package config

import (
	"sync"

	"github.com/spf13/viper"
)

var mu sync.Mutex

// LoadConfig initiates of config load
func LoadConfig(configPath string) error {
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	LoadApp()
	LoadDB()
	LoadMinio()
	LoadRabbitMQ()

	return nil
}
