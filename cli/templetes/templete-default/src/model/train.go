package model

import (
	"fmt"

	"github.com/RuiHirano/simframe/app"
)

type Train struct {
	*app.Agent
}

func NewTrain(id string, position *app.Position) *Train {

	train := &Train{
		app.NewAgent(id, "train", position),
	}

	return train
}

func (train *Train) Step() {
	fmt.Printf("Train %s step\n", train.ID)
}
