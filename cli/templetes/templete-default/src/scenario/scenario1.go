package scenario

import (

	"github.com/RuiHirano/simframe/model"
	"github.com/RuiHirano/simframe/scenario"
	mymodel "github.com/RuiHirano/simframe/src/model"
	"github.com/RuiHirano/simframe/util"
)

type Scenario1 struct {
	*scenario.Scenario
}

func NewScenario1() *Scenario1 {

	agents := []model.IAgent{
		mymodel.NewTrain("1", &util.Position{X: 2, Y: 4}),
		mymodel.NewTrain("2", &util.Position{X: 3, Y: 3}),
		mymodel.NewCar("3", &util.Position{X: 1, Y: 5}),
	}
	area := scenario.NewArea()

	clock := scenario.NewClock()

	scenario := &Scenario1{
		scenario.NewScenario(agents, area, clock),
	}

	return scenario
}
