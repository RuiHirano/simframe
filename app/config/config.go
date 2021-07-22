package config

import (
	"github.com/RuiHirano/simframe/api"
)

type IConfig interface {
}

type Config struct {
	*api.Config
}

type ConfigParams struct {
	ServerNum int64
}

func NewConfig(param *ConfigParams) *Config {

	config := &Config{
		&api.Config{
			ServerNum: param.ServerNum,
		},
	}

	return config
}