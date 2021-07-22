package api


func NewApp(sns []*Scenario, conf *Config) *App {

	app := &App{
		Scenarios: sns,
		Config: conf,
	}

	return app
}