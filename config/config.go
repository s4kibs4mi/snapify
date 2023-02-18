package config

import (
	"github.com/spf13/viper"
	"os"
	"time"
)

func LoadConfig() (*AppCfg, error) {
	v := viper.New()
	v.SetConfigFile(os.Getenv("CONFIG_FILE")) // To set config file dynamically using environment variable

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return LoadConfigWithViper(v)
}

func LoadConfigWithViper(v *viper.Viper) (*AppCfg, error) {
	return &AppCfg{
		Base:              v.GetString("app.host"),
		Port:              v.GetInt("app.port"),
		LogLevel:          LogLevel(v.GetString("app.log_level")),
		ChromeHeadlessUrl: v.GetString("app.chrome_headless_url"),
		RedisAddr:         v.GetString("app.redis_addr"),
		RedisUsername:     v.GetString("app.redis_username"),
		RedisPassword:     v.GetString("app.redis_password"),
		RedisQueueName:    v.GetString("app.redis_queue_name"),
		DBCfg: DBCfg{
			Name:                  v.GetString("database.name"),
			Username:              v.GetString("database.username"),
			Password:              v.GetString("database.password"),
			Host:                  v.GetString("database.host"),
			Port:                  v.GetInt("database.port"),
			MaxOpenConnections:    v.GetInt("database.max_open_connections"),
			MaxIdleConnections:    v.GetInt("database.max_idle_connections"),
			MaxConnectionLifetime: v.GetDuration("database.max_connection_lifetime") * time.Minute,
		},
		BlobStorageCfg: BlobStorageCfg{
			BaseURL:      v.GetString("minio.base_url"),
			Key:          v.GetString("minio.key"),
			Secret:       v.GetString("minio.secret"),
			Bucket:       v.GetString("minio.bucket"),
			Location:     v.GetString("minio.location"),
			SignDuration: v.GetDuration("minio.sign_duration") * time.Minute,
			IsSecure:     v.GetBool("minio.is_secure"),
		},
	}, nil
}
