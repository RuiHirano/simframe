package app

import (
	"fmt"
)

type IAgent interface {
	Status()
	Step()
}

type Agent struct {
	ID string
	Name string
	Position *Position
}

func NewAgent(id string, name string, position *Position) *Agent {

	agent := &Agent{
		ID: id,
		Name: name,
		Position: position,
	}

	return agent
}

func (agent *Agent) Status() {
	fmt.Printf("Status %s %v\n", agent.ID, agent.Position)
}