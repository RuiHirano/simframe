package train

import (
	"fmt"

	"github.com/RuiHirano/simframe/model"
)

type Train struct {
	*model.Agent
}

func NewTrain(id string) *Train {

	train := &Train{
		model.NewAgent(id),
	}

	return train
}

func (train *Train) Step() {
	fmt.Printf("Train %s step\n", train.ID)
}
