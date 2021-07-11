package train

import (
	"fmt"
	"github.com/RuiHirano/simframe/util"

	"github.com/RuiHirano/simframe/model"
)

type Train struct {
	*model.Agent
}

func NewTrain(id string, position *util.Position) *Train {

	train := &Train{
		model.NewAgent(id, position),
	}

	return train
}

func (train *Train) Step() {
	fmt.Printf("Train %s step\n", train.ID)
}
