package app

type IScenario interface {
	GetArea() IArea
	GetClock() IClock
	GetAgents() []IAgent
}

type Scenario struct {
	Area IArea
	Clock IClock
	Agents []IAgent
}

func NewScenario(agents []IAgent, area IArea, clock IClock) *Scenario {

	scenario := &Scenario{
		Area: area,
		Clock: clock,
		Agents: agents,
	}

	return scenario
}

func (sn *Scenario) GetArea() IArea{
	return sn.Area
}

func (sn *Scenario) GetClock() IClock{
	return sn.Clock
}

func (sn *Scenario) GetAgents() []IAgent{
	return sn.Agents
}