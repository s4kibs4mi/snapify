package config

import (
	"github.com/spf13/viper"
	"time"
)

type MinioCfg struct {
	BaseURL      string
	ServeURL     string
	Key          string
	Secret       string
	Location     string
	Bucket       string
	SignDuration time.Duration
}

var minio MinioCfg

func LoadMinio() {
	mu.Lock()
	defer mu.Unlock()

	minio = MinioCfg{
		BaseURL:      viper.GetString("minio.base_url"),
		ServeURL:     viper.GetString("minio.serve_url"),
		Key:          viper.GetString("minio.key"),
		Secret:       viper.GetString("minio.secret"),
		Bucket:       viper.GetString("minio.bucket"),
		Location:     viper.GetString("minio.location"),
		SignDuration: viper.GetDuration("minio.sign_duration") * time.Minute,
	}
}

func Minio() MinioCfg {
	return minio
}
