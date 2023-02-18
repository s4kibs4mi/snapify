package config

import (
	"time"
)

type DBCfg struct {
	Host                  string
	Port                  int
	Username              string
	Password              string
	Name                  string
	MaxOpenConnections    int
	MaxIdleConnections    int
	MaxConnectionLifetime time.Duration
}
