package main

import (
	"github.com/RuiHirano/simframe/app"
	"github.com/RuiHirano/simframe/app/config"
	"github.com/RuiHirano/simframe/scenario"
	myscenario "github.com/RuiHirano/simframe/src/scenario"
)

type App struct{
	app.App
}

func (ap *App) Scenarios() []scenario.IScenario{
	sn := myscenario.NewScenario1()
	return []scenario.IScenario{sn}
}

func (ap *App) Config() config.IConfig{
	config := config.NewConfig(&config.ConfigParams{
		ServerNum: 1,
	})
	return config
}