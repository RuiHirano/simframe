package myapp

import (
	myscenario "scenario"

	"github.com/RuiHirano/simframe/app/config"
	"github.com/RuiHirano/simframe/app/scenario"
)

func Scenarios() []scenario.IScenario{
	sn := myscenario.NewScenario1()
	return []scenario.IScenario{sn}
}

func Config() config.IConfig{
	config := config.NewConfig(&config.ConfigParams{
		ServerNum: 1,
	})
	return config
}