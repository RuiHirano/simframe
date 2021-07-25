package app

type IApp interface {
	GetScenarios() []IScenario
	GetConfig() IConfig
}

type App struct {
	Scenarios []IScenario
	Config IConfig
}

func NewApp(sns []IScenario, conf IConfig) IApp {

	app := &App{
		Scenarios: sns,
		Config: conf,
	}

	return app
}

func (app *App) GetScenarios() []IScenario{
	return app.Scenarios
}

func (app *App) GetConfig() IConfig{
	return app.Config
}