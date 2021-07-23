package app

type IScenario interface {
}

type Scenario struct {
	Area IArea
	Clock IClock
	Agents []IAgent
}

func NewScenario(agents []IAgent, area IArea, clock IClock) IScenario {

	scenario := &Scenario{
		Area: area,
		Clock: clock,
		Agents: agents,
	}

	return scenario
}