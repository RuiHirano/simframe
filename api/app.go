package api

func NewApp(scenarios []*Scenario, conf *Config) *App {

	app := &App{
		Scenarios: scenarios,
		Config: conf,
	}

	return app
}
