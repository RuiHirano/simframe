package app

import (
	"github.com/RuiHirano/simframe/api"
	"github.com/RuiHirano/simframe/app/config"
	"github.com/RuiHirano/simframe/app/scenario"
)

type IApp interface {
	GetConfig() config.IConfig
	GetScenarios() []scenario.IScenario
}

type App struct {
	*api.App
}

func NewApp(scenarios []scenario.IScenario, conf config.IConfig) *App {

	app := &App{
		&api.App{
			Scenarios: scenarios,
			Config: conf,
		},
	}

	return app
}

func (app *App) GetConfig() config.IConfig{
	return app.Config
}

func (app *App) GetScenarios() []scenario.IScenario{
	return app.Scenarios
}