package model

import (
	"fmt"

	"github.com/RuiHirano/simframe/app"
)

type Car struct {
	*app.Agent
}

func NewCar(id string, position *app.Position) app.IAgent {

	car := &Car{
		&app.Agent{
			ID: id, 
			Name: "CAR",
			Position: position,
		},
	}

	return car
}

func (car *Car) Step() {
	fmt.Printf("Car %s step : %v\n", car.ID, car.Position)
	car.Position = car.NextPosition()
}

func (car *Car) NextPosition() *app.Position{
	nextPosition := &app.Position{X: car.Position.X+1, Y: car.Position.Y+1}
	return nextPosition
}