package engine

import (
	"github.com/RuiHirano/simframe/model"
)

type IEngine interface {
	Run()
}

type Engine struct {
	ID string
	Agents []model.IAgent
	Clock IClock
}

func NewEngine(agents []model.IAgent) *Engine {

	engine := &Engine{
		ID: "0",
		Agents: agents,
		Clock: NewClock(),
	}

	return engine
}

func (engine *Engine) Run() {
	for i := 0; i < 10; i++ {
		for _, agent := range engine.Agents{
			agent.Step()
			agent.Status()
		}
		engine.Clock.Forward()
		engine.Clock.GetTimestamp()
	}
}