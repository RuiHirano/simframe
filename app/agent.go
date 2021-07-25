package app

import (
	"fmt"
)

type IAgent interface {
	Step()
	GetID() string
	GetName() string
	GetPosition() *Position
}

type Agent struct {
	ID string
	Name string
	Position *Position
}

func NewAgent(id string, name string, position *Position) IAgent {

	agent := &Agent{
		ID: id,
		Name: name,
		Position: position,
	}

	return agent
}

func (agent *Agent) GetID() string{
	return agent.ID
}

func (agent *Agent) GetName() string{
	return agent.Name
}

func (agent *Agent) GetPosition() *Position{
	return agent.Position
}

func (agent *Agent) Step(){
	fmt.Printf("Agent %s step : %v\n", agent.ID, agent.Position)
}