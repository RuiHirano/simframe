package github.com/RuiHirano/simframe/app/config

type IConfig interface {
}

type Config struct {
	ServerNum int64
}

type ConfigParams struct {
	ServerNum int64
}

func NewConfig(param *ConfigParams) *Config {

	config := &Config{
		ServerNum: param.ServerNum,
	}

	return config
}