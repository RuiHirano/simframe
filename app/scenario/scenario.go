package github.com/RuiHirano/simframe/app/scenario

import (
	"github.com/RuiHirano/simframe/app/model"
)

type IScenario interface {
	Run()
	Step()
}

type Scenario struct {
	ID string
	Agents []model.IAgent
	Clock IClock
	Area IArea
}

func NewScenario(agents []model.IAgent, area IArea, clock IClock) *Scenario {

	scenario := &Scenario{
		ID: "0",
		Agents: agents,
		Clock: clock,
		Area: area,
	}

	return scenario
}

// TODO: Remove Run
func (scenario *Scenario) Run() {
	scenario.Area.GetSpace()
	for i := 0; i < 10; i++ {
		for _, agent := range scenario.Agents{
			agent.Step()
			agent.Status()
		}
		scenario.Clock.Forward()
		scenario.Clock.GetTimestamp()
	}
}

func (scenario *Scenario) Step(nbAgents []model.IAgent) {
	for _, agent := range scenario.Agents{
		agent.Step()
		agent.Status()
	}
	scenario.Clock.Forward()
	scenario.Clock.GetTimestamp()
}