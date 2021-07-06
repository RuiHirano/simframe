package model

import (
	"fmt"

	"github.com/RuiHirano/simframe/util"
)

type IAgent interface {
	Status()
	Step()
}

type Agent struct {
	ID string
	Position *util.Position
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