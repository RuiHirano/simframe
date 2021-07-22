package api

type ConfigParams struct {
	ServerNum uint64
}

func NewConfig(param *ConfigParams) *Config {

	config := &Config{
		ServerNum: param.ServerNum,
	}

	return config
}