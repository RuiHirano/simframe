package train

import (
	"fmt"

	"github.com/RuiHirano/simframe/util"

	"github.com/RuiHirano/simframe/app/model"
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
	fmt.Printf("Car %s step : %v\n", car.ID, car.Position)
	car.Position = car.NextPosition()
}

func (car *Car) NextPosition() *util.Position{
	nextPosition := &util.Position{X: car.Position.X+1, Y: car.Position.Y+1}
	return nextPosition
}