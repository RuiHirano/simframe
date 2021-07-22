package api

import (
	"fmt"
)

func NewAgent(id string, position *Position) *Agent {

	agent := &Agent{
		Id: id,
		Position: position,
	}

	return agent
}

func (agent *Agent) Status() {
	fmt.Printf("Status %s %v\n", agent.Id, agent.Position)
}

func (agent *Agent) Step() {
	fmt.Printf("Status %s %v\n", agent.Id, agent.Position)
}
