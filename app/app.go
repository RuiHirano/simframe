package app

type IApp interface {
	GetScenarios() []IScenario
}

type App struct {
	Scenarios []IScenario
}

func NewApp(sns []IScenario) *App {

	app := &App{
		Scenarios: sns,
	}

	return app
}

func (app *App) GetScenarios() []IScenario{
	return app.Scenarios
}