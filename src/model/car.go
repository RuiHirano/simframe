package train

import (
	"fmt"

	"github.com/RuiHirano/simframe/util"

	"github.com/RuiHirano/simframe/model"
)

type Car struct {
	*model.Agent
}

func NewCar(id string, position *util.Position) *Car {

	car := &Car{
		model.NewAgent(id, position),
	}

	return car
}

func (car *Car) Step() {
	fmt.Printf("Car %s step\n", car.ID)
}
