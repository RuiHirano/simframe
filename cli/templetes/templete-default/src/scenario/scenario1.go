package scenario

import (
	mymodel "model"

	"github.com/RuiHirano/simframe/app"
)

type Scenario1 struct {
	app.IScenario
}

func NewScenario1() app.IScenario {

	agents := []app.IAgent{
		mymodel.NewTrain("1", &app.Position{X: 2, Y: 4}),
		mymodel.NewTrain("2", &app.Position{X: 3, Y: 3}),
		mymodel.NewCar("3", &app.Position{X: 1, Y: 5}),
	}
	area := app.NewArea(&app.Space{MinX: 0, MaxX:10, MinY:0, MaxY:10})

	clock := app.NewClock(0)

	scenario := &Scenario1{
		app.NewScenario(agents, area, clock),
	}

	scenario.RegisterModel(
		[]app.Model{
			mymodel.NewTrain,
			mymodel.NewCar,
		},
	)

	return scenario
}
