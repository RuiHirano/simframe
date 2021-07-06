package model

import (
	"fmt"
)

type IAgent interface {
	Status()
	Step()
}

type Agent struct {
	ID string
}

func NewAgent(id string) *Agent {

	agent := &Agent{
		ID: id,
	}

	return agent
}

func (agent *Agent) Status() {
	fmt.Printf("Status %s\n", agent.ID)
}