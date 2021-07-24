package engine

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/RuiHirano/simframe/app"
)

type IEngine interface {
	Run(runType string)
}

type Engine struct {
	App app.IApp
}

func NewEngine(ap app.IApp) *Engine {

	engine := &Engine{
		App: ap,
	}

	return engine
}

func (engine *Engine) Run(runType string) {
	//unType := "WORKER"
	switch runType {
	case "ENGINE":
		svc := NewEngineService(10000, engine.RegisterSimulatorHandler)
		go svc.Serve()
		generator := NewResourceGenerator(engine.App)
		for i := 0; i < 4; i++ {
			generator.Apply(strconv.Itoa(i), 9000+i)
		}

	case "SIMULATOR":
		sim := NewSimulator()
		svc := NewSimulatorService(9000, sim.RunSimulatorHandler, sim.GetNeighborAgentsHandler)
		go svc.Serve()
		sim.Run()
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
}

func (engine *Engine) RegisterSimulatorHandler() (app.IArea, app.IClock, []app.IAgent){
	fmt.Printf("Register simulator\n")
	return app.NewArea(), app.NewClock(), []app.IAgent{}
}
