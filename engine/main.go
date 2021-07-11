package engine

import (
	"fmt"

	"github.com/RuiHirano/simframe/config"
	"github.com/RuiHirano/simframe/scenario"
)

type IEngine interface {
	Run()
}

type Engine struct {
	ID string 
	Scenarios []scenario.IScenario
	Config config.IConfig
}

func NewEngine(scenarios []scenario.IScenario, conf config.IConfig) *Engine {

	engine := &Engine{
		ID: "0",
		Scenarios: scenarios,
		Config: conf,
	}

	return engine
}

func (engine *Engine) Run() {
	fmt.Printf("Run Engine\n")
	for _, scenario := range engine.Scenarios{
		scenario.Run()
		// Run scenario in docker container
	}
}