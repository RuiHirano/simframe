package app

import (
	"github.com/RuiHirano/simframe/app/config"
	"github.com/RuiHirano/simframe/app/scenario"
)

type IApp interface {
	GetConfig() config.IConfig
	GetScenarios() []scenario.IScenario
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

func (app *App) GetConfig() config.IConfig{
	return app.Config
}

func (app *App) GetScenarios() []scenario.IScenario{
	return app.Scenarios
}