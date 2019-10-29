package config

import (
	"github.com/spf13/viper"
)

type LogLevel string

const (
	Debug    LogLevel = "debug"
	Info     LogLevel = "info"
	Warning  LogLevel = "warning"
	Critical LogLevel = "critical"
)

type Application struct {
	Base              string
	Port              int
	LogLevel          LogLevel
	ChromeHeadlessUrl string
}

var app Application

func App() *Application {
	return &app
}

func LoadApp() {
	mu.Lock()
	defer mu.Unlock()

	app = Application{
		Base:              viper.GetString("app.host"),
		Port:              viper.GetInt("app.port"),
		LogLevel:          LogLevel(viper.GetString("app.log_level")),
		ChromeHeadlessUrl: viper.GetString("app.chrome_headless_url"),
	}
}
