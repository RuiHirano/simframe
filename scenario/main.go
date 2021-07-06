package scenario

import (
	"github.com/RuiHirano/simframe/model"
)

type IScenario interface {
	Run()
}

type Scenario struct {
	ID string
	Agents []model.IAgent
	Clock IClock
	Area IArea
}

func NewScenario(agents []model.IAgent) *Scenario {

	scenario := &Scenario{
		ID: "0",
		Agents: agents,
		Clock: NewClock(),
		Area: NewArea(),
	}

	return scenario
}

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