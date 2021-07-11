package main

import (
	"github.com/RuiHirano/simframe/config"
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

	area := scenario.NewArea()

	clock := scenario.NewClock()

	sn := scenario.NewScenario(agents, area, clock)

	config := config.NewConfig(&config.ConfigParams{
		ServerNum: 1,
	})

	en := engine.NewEngine([]scenario.IScenario{sn}, config)
	en.Run()
}