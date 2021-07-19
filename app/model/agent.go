package github.com/RuiHirano/simframe/app/model

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

func NewAgent(id string, position *util.Position) *Agent {

	agent := &Agent{
		ID: id,
		Position: position,
	}

	return agent
}

func (agent *Agent) Status() {
	fmt.Printf("Status %s %v\n", agent.ID, agent.Position)
}