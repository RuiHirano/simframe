package app

import (
	"github.com/RuiHirano/simframe/app/scenario"
	"github.com/RuiHirano/simframe/app/config"
)

type IApp interface {
	Scenarios() []scenario.IScenario
	Config() config.IConfig
}

type App struct {
	Scenarios []scenario.IScenario
	Config config.IConfig
}

func NewApp(scenarios []scenario.IScenario, conf config.IConfig) *App {

	app := &App{
		Scenarios: scenarios,
		Config: conf,
	}

	return app
}