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
	Scenarios []scenario.IScenario,
	Config config.IConfig
}
