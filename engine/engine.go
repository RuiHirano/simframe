package engine

import (
	"fmt"
	//"strconv"
	"sync"

	"github.com/RuiHirano/simframe/app"
	"github.com/fatih/color"
)

type IEngine interface {
	Run(runType string)
	GetApp() app.IApp
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
		engine.Start()
		/*generator := NewResourceGenerator(engine.App)
		for i := 0; i < 4; i++ {
			generator.Apply(strconv.Itoa(i), 9000+i)
		}*/

	case "SIMULATOR":
		sim := NewSimulator(engine.App)
		sim.Run()
		
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func (engine *Engine) GetApp() app.IApp{
	return engine.App
}

func (engine *Engine) SetUp(){
	// service登録
	color.Green("Register engine service...\n")
	svc := NewEngineService(10000, engine.RegisterSimulatorHandler)
	go svc.Serve()
}

func (engine *Engine) Start() {
	fmt.Printf("Run Engine\n")
	engine.SetUp()
}

func (engine *Engine) RegisterSimulatorHandler() (app.IArea, app.IClock, []app.IAgent, []*Neighbor, int){
	fmt.Printf("Register simulator\n", engine.App.GetScenarios()[0].GetAgents()[0].GetName())
	return engine.App.GetScenarios()[0].GetArea(), engine.App.GetScenarios()[0].GetClock(), engine.App.GetScenarios()[0].GetAgents(), []*Neighbor{} , 9001
}
