package main

import (
	"github.com/RuiHirano/simframe/engine"
	"github.com/RuiHirano/simframe/model"
	"github.com/RuiHirano/simframe/scenario"
	mymodel "github.com/RuiHirano/simframe/src/model"
	"github.com/RuiHirano/simframe/util"
)

func main() {
    agents := []model.IAgent{
		mymodel.NewTrain("1", &util.Position{X: 2, Y: 4}),
		mymodel.NewTrain("2", &util.Position{X: 3, Y: 3}),
		mymodel.NewCar("3", &util.Position{X: 1, Y: 5}),
	}

	var sn scenario.IScenario
	sn = scenario.NewScenario(agents)

	//var sn2 scenario.IScenario
	//sn2 = scenario.NewScenario(agents)

	var en engine.IEngine
	en = engine.NewEngine([]scenario.IScenario{sn})
	en.Run()
}