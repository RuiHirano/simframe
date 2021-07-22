package app

import (
	"github.com/RuiHirano/simframe/api"
)

type IScenario interface {
	Run()
	Step(agents []IAgent)
}

type Scenario struct {
	*api.Scenario
}

func NewScenario(agents []IAgent, area IArea, clock IClock) *Scenario {

	scenario := &Scenario{
		&api.Sceanrio{
			ID: "0",
			Agents: agents,
			Clock: clock,
			Area: area,
		},
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

func (scenario *Scenario) Step(nbAgents []IAgent) {
	for _, agent := range scenario.Agents{
		agent.Step()
		agent.Status()
	}
	scenario.Clock.Forward()
	scenario.Clock.GetTimestamp()
}