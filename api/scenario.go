package api

func NewScenario(agents []*Agent) *Scenario {
	scenario := &Scenario{
		Id: "0",
		Agents: agents,
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

func (scenario *Scenario) Step(nbAgents []*Agent) {
	for _, agent := range scenario.Agents{
		agent.Step()
		agent.Status()
	}
	scenario.Clock.Forward()
	scenario.Clock.GetTimestamp()
}