package engine

import (
	"fmt"
	"runtime"

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
	SimulatorManager *SimulatorManager
}

func NewEngine(ap app.IApp) *Engine {

	engine := &Engine{
		App: ap,
		SimulatorManager: NewSimulatorManager(ap),
	}

	return engine
}

func (engine *Engine) Run(runType string, id string, port int) {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	//unType := "WORKER"
	switch runType {
	case "ENGINE":
		engine.Start()

	case "SIMULATOR":
		sim := NewSimulator(engine.App, id, port)
		sim.Start()
		
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

	sm := engine.SimulatorManager
	sm.CalculateSimulators()
	sm.ApplySimulators()
	sm.WaitSimulators()
	sm.SetUpSimulators()
	sm.RunSimulators()

	
}

func (engine *Engine) RegisterSimulatorHandler(id string) int{
	fmt.Printf("Register simulator %v\n", id)
	engine.SimulatorManager.RegisterSimulator(id)
	port := 9000
	return port
}

type ISimulatorManager interface {
	CalculateSimulators()
	ApplySimulators()
	WaitSimulators()
	SetUpSimulators()
	RunSimulators()
}

type SimulatorManager struct {
	App app.IApp
	Simulators []*SimulatorData
	Syncronizer *Syncronizer
}

type SimulatorData struct{
	ID string
	Port int
	Neighbors []*Neighbor
	API *SimulatorAPI
}

func NewSimulatorManager(ap app.IApp) *SimulatorManager {

	sm := &SimulatorManager{
		App: ap,
		Syncronizer: NewSyncronizer(),
	}

	return sm
}

func (sm *SimulatorManager) CalculateSimulators() {
	color.Green("Calculating simulators...\n")
	sm.Simulators = []*SimulatorData{
		&SimulatorData{
			ID: "simulator1", 
			Port: 9000, 
			Neighbors: []*Neighbor{
				&Neighbor{ID: "simulator2", Port: 9001},
			},
		},
		&SimulatorData{
			ID: "simulator2", 
			Port: 9001, 
			Neighbors: []*Neighbor{
				&Neighbor{ID: "simulator1", Port: 9000},
			},
		},
	}
	
}

func (sm *SimulatorManager) ApplySimulators() {
	color.Green("Applying simulators...\n")
	/*generator := NewResourceGenerator(engine.App)
	for i := 0; i < 4; i++ {
		generator.Apply(strconv.Itoa(i), 9000+i)
	}*/
}

func (sm *SimulatorManager) WaitSimulators() {
	color.Green("Waiting simulators...\n")
	simIds := []string{}
	for _, sim := range sm.Simulators{
		simIds = append(simIds, sim.ID)
	}
	sm.Syncronizer.RegisterIds("WAIT_SIMULATORS", simIds)
	sm.Syncronizer.Wait("WAIT_SIMULATORS")
	
}

func (sm *SimulatorManager) RegisterSimulator(id string) {
	color.Green("Registered simulator [id:%s]\n", id)
	for i, sim := range sm.Simulators{
		if sim.ID == id{
			color.Green("Connecting to simulator [id:%s, port:%d]\n", id, sim.Port)
			sm.Simulators[i].API = NewSimulatorAPI(sim.Port)
			sm.Simulators[i].API.SetUp()
		}
	}
	sm.Syncronizer.AddId("WAIT_SIMULATORS", id)
}

func (sm *SimulatorManager) SetUpSimulators() {
	color.Green("Setting up simulators...\n")
	for _, sim := range sm.Simulators{
		sim.API.SetUpSimulator(
			sm.App.GetScenarios()[0].GetArea(), 
			sm.App.GetScenarios()[0].GetClock(), 
			sm.App.GetScenarios()[0].GetAgents(), 
			sim.Neighbors,
		)
	}
}

func (sm *SimulatorManager) RunSimulators() {
	color.Green("Running simulators\n")
	for _, sim := range sm.Simulators{
		sim.API.RunSimulator("engine")
	}
}