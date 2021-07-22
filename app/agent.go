package app

import (
	"fmt"

	"github.com/RuiHirano/simframe/api"
)

type IAgent interface {
	Status()
	Step()
}

type Agent struct {
	*api.Agent
}

func NewAgent(id string, position *api.Position) *Agent {

	agent := &Agent{
		&api.Agent{
			Id: id,
			Position: position,
		},
	}

	return agent
}

func (agent *Agent) Status() {
	fmt.Printf("Status %s %v\n", agent.Id, agent.Position)
}