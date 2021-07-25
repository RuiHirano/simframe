package engine

import (
	"fmt"

	"github.com/RuiHirano/simframe/app"
	"github.com/fatih/color"
)

type Neighbor struct{
	ID string
	Port int
	API *SimulatorAPI
}

type ISimulator interface {
	Run()
	GetID() string
	GetPort() int
	GetNeighbors() []*Neighbor
	GetClock() app.IClock
	GetAgents() []app.IAgent
	GetArea() app.IArea
}

type Simulator struct {
	ID string 
	App app.IApp
	EngineAPI *EngineAPI
	SimulatorAPI *SimulatorAPI 
	Port int
	Neighbors []*Neighbor
	Clock app.IClock
	Agents []app.IAgent
	Area app.IArea
}

func NewSimulator(ap app.IApp) *Simulator {

	sim := &Simulator{
		ID: "0",
		App: ap,
		EngineAPI: NewEngineAPI(10000, ap),
	}
	return sim
}

func (sim *Simulator) SetUp(){
	// engineから必要なareaを取得する
	color.Green("Register simulator to engine...\n")
	sim.Area, sim.Clock, sim.Agents, sim.Neighbors, sim.Port = sim.EngineAPI.RegisterSimulator(sim.ID, sim)
	fmt.Printf("test", sim.Neighbors, sim.Port)
	// service登録
	color.Green("Register simulator service...\n")
	svc := NewSimulatorService(sim.Port, sim.RunSimulatorHandler, sim.GetNeighborAgentsHandler)
	go svc.Serve()
}

func (sim *Simulator) GetID() string{
	return sim.ID
}
func (sim *Simulator) GetPort() int{
	return sim.Port
}

func (sim *Simulator) GetNeighbors() []*Neighbor{
	return sim.Neighbors
}

func (sim *Simulator) GetClock() app.IClock{
	return sim.Clock
}

func (sim *Simulator) GetAgents() []app.IAgent{
	return sim.Agents
}

func (sim *Simulator) GetArea() app.IArea{
	return sim.Area
}


func (sim *Simulator) SyncNeighborArea() []app.IAgent{
	// send Ack

	// wait Ack from all neighbors
	//syn := NewSyncronizer()
	// send sync request
	//for _, nei := range sim.Neighbors{
		//syn.RegisterId(nei.GetID())
		//sim.Agents = sim.SimulatorAPI.GetNeighborAgents(sim.ID, []app.IAgent{})
	//}
	//syn.Wait()

	return []app.IAgent{}
}

func (sim *Simulator) UpdateAgents(agents []app.IAgent){}
func (sim *Simulator) CalculateAgents(){
	for _, agent := range sim.Agents{
		agent.Step()
	}
}
func (sim *Simulator) ForwardClock(){
	sim.Clock.Forward()
}
func (sim *Simulator) Step(){
	fmt.Printf("Step Time: %d\n", sim.Clock.GetTimestamp())
	// 隣接エリアとのClock同期をとる
	agents := sim.SyncNeighborArea()

	// 隣接エリアのエージェントを取得
	sim.UpdateAgents(agents)

	// エージェント計算を行う
	sim.CalculateAgents()

	// Clockを進める
	sim.ForwardClock()
}

func (sim *Simulator) Run() {
	fmt.Printf("Run Simulator\n")
	sim.SetUp()
	for i := 0; i < 10; i++ {
		sim.Step()
	}
	fmt.Printf("Simulator Finished\n")
}

func (sim *Simulator) RunSimulatorHandler(){
	fmt.Printf("Run simulator\n")
	sim.Run()
}

func (sim *Simulator) GetNeighborAgentsHandler(){
	fmt.Printf("Get Neighbor Agents\n")
	
}