package config

type LogLevel string

const (
	Debug    LogLevel = "debug"
	Info     LogLevel = "info"
	Warning  LogLevel = "warning"
	Critical LogLevel = "critical"
)

type AppCfg struct {
	Base               string
	Port               int
	LogLevel           LogLevel
	RedisAddr          string
	RedisPassword      string
	RedisUsername      string
	RedisQueueName     string
	HeadlessBrowserUrl string
	AuthToken          string
	DBCfg              DBCfg
	BlobStorageCfg     BlobStorageCfg
}
