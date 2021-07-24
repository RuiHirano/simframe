package scenario

import (
	mymodel "model"

	"github.com/RuiHirano/simframe/app"
)

type Scenario1 struct {
	*app.Scenario
}

func NewScenario1() *Scenario1 {

	agents := []app.IAgent{
		mymodel.NewTrain("1", &app.Position{X: 2, Y: 4}),
		mymodel.NewTrain("2", &app.Position{X: 3, Y: 3}),
		mymodel.NewCar("3", &app.Position{X: 1, Y: 5}),
	}
	area := app.NewArea()

	clock := app.NewClock()

	scenario := &Scenario1{
		app.NewScenario(agents, area, clock),
	}

	return scenario
}
