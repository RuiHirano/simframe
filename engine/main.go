package engine

import (
	"fmt"

	"github.com/RuiHirano/simframe/scenario"
)

type IEngine interface {
	Run()
}

type Engine struct {
	ID string
	Scenarios []scenario.IScenario
}

func NewEngine(scenarios []scenario.IScenario) *Engine {

	engine := &Engine{
		ID: "0",
		Scenarios: scenarios,
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