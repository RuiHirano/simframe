package api

import (
	"fmt"
)

type IAgent interface {
	Status()
	Step()
	GetGrpcAgent() *Agent
}

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

func (agent *Agent) GetGrpcAgent() *Agent{
	return agent
}

