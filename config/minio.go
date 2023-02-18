package config

import (
	"time"
)

type BlobStorageCfg struct {
	BaseURL      string
	Key          string
	Secret       string
	Location     string
	Bucket       string
	SignDuration time.Duration
	IsSecure     bool
}
