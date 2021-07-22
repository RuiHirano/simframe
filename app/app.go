package app

import (
	"github.com/RuiHirano/simframe/api"
)

type IApp interface {
	GetConfig() IConfig
	GetScenarios() []IScenario
}

type App struct {
	*api.App
}

func NewApp(scenarios []IScenario, conf IConfig) *App {

	app := &App{
		&api.App{
			Scenarios: scenarios,
			Config: conf,
		},
	}

	return app
}

func (app *App) GetConfig() IConfig{
	return app.Config
}

func (app *App) GetScenarios() []IScenario{
	return app.Scenarios
}