package model

import (
	"fmt"

	"github.com/RuiHirano/simframe/app"
)

type Train struct {
	*app.Agent
}

func NewTrain(id string, position *app.Position) app.IAgent {

	train := &Train{
		&app.Agent{
			ID: id, 
			Name: "TRAIN",
			Position: position,
		},
	}

	return train
}

func (train *Train) Step() {
	fmt.Printf("Train %s step : %v\n", train.ID, train.Position)
	train.Position = train.NextPosition()
}

func (train *Train) NextPosition() *app.Position{
	nextPosition := &app.Position{X: train.Position.X+10, Y: train.Position.Y+15}
	return nextPosition
}