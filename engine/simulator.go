package engine

import (
	"fmt"
	"os"

	"github.com/RuiHirano/simframe/app"
	"github.com/fatih/color"
	"github.com/google/uuid"
)

func GetUUID()string{
	u, _ := uuid.NewRandom()
	return u.String()
}

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
	Syncronizer *Syncronizer
	Port int
	Neighbors []*Neighbor
	Clock app.IClock
	Agents []app.IAgent
	Area app.IArea
}

func NewSimulator(ap app.IApp, id string, port int) *Simulator {

	sim := &Simulator{
		ID: id,
		Port: port,
		App: ap,
		EngineAPI: NewEngineAPI(10000, ap),
		Syncronizer: NewSyncronizer(),
	}
	return sim
}

func (sim *Simulator) Start() {
	svc := NewSimulatorService(sim.App, sim.Port, sim.SetUp, sim.Run, sim.GetNeighborAgents)
	go svc.Serve()
	sim.EngineAPI.SetUp()
	sim.RegisterSimulator()
}

func (sim *Simulator) RegisterSimulator(){
	err, _ := sim.EngineAPI.RegisterSimulator(sim.ID)
	if err != nil{
		color.Red("Cannot connect to engine.\n")
		os.Exit(1)
	}
	color.Green("Success Register to engine.\n")
}

func (sim *Simulator) SetUp(area app.IArea, clock app.IClock, agents []app.IAgent, neighbors []*Neighbor){
	fmt.Printf("SetUp simulator\n")
	sim.Area, sim.Clock, sim.Agents, sim.Neighbors = area, clock, agents, neighbors
	// service登録
	color.Green("Register simulator service...\n", sim.Neighbors[0].ID)

	// neighborsに登録
	neiIds := []string{}
	for i, nei := range sim.Neighbors{
		sim.Neighbors[i].API = NewSimulatorAPI(nei.Port)
		sim.Neighbors[i].API.SetUp()
		neiIds = append(neiIds, nei.ID)
	}
	color.Green("Register neighbors... %v\n", neiIds)
	sim.Syncronizer.RegisterIds("WAIT_NEIGHBORS", neiIds)
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
	// neighborsとsync
	for i, _ := range sim.Neighbors{
		sim.Neighbors[i].API.GetNeighborAgents(sim.ID, sim.Agents)
	}

	// wait Ack from all neighbors
	color.Green("Waiting neighbors...\n")
	/*neiIds := []string{}
	for _, nei := range sim.Neighbors{
		neiIds = append(neiIds, nei.ID)
	}*/
	sim.Syncronizer.Wait("WAIT_NEIGHBORS")

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
	color.Yellow("Step Time: %d\n", sim.Clock.GetTimestamp())
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
	//sim.SetUp()
	go func(){
		for i := 0; i < 10000; i++ {
			sim.Step()
		}
		color.Green("Simulator Finished\n")
	}()
}


func (sim *Simulator) GetNeighborAgents(id string, agents []app.IAgent){
	//fmt.Printf("Get Neighbor Agents from %s\n", id)
	sim.Syncronizer.AddId("WAIT_NEIGHBORS", id)
}