package main

import (
	"github.com/RuiHirano/simframe/engine"
	"github.com/RuiHirano/simframe/model"
	mymodel "github.com/RuiHirano/simframe/src/model"
)

func main() {
    agents := []model.IAgent{
		mymodel.NewTrain("1"),
		mymodel.NewTrain("2"),
		mymodel.NewCar("3"),
	}

	var en engine.IEngine
	en = engine.NewEngine(agents)
	en.Run()
}