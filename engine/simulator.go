package engine

import (
	"fmt"

	"github.com/RuiHirano/simframe/app"
)

type ISimulator interface {
	Run()
	GetID() string
	GetServerAddress() string
	GetNeighbors() []ISimulator
	GetClock() app.IClock
	GetAgents() []app.IAgent
	GetArea() app.IArea
}

type Simulator struct {
	ID string 
	ServerAddress string
	Neighbors []ISimulator
	Clock app.IClock
	Agents []app.IAgent
	Area app.IArea
}

func NewSimulator() *Simulator {

	Simulator := &Simulator{
		ID: "0",
	}

	return Simulator
}



func (sim *Simulator) GetID() string{
	return sim.ID
}
func (sim *Simulator) GetServerAddress() string{
	return sim.ServerAddress
}

func (sim *Simulator) GetNeighbors() []ISimulator{
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

func (sim *Simulator) GetNeighborData() []app.IAgent{
	agents := []app.IAgent{}
	return agents
}

func (sim *Simulator) StepAgents(){}
func (sim *Simulator) ClockForward(){}
func (sim *Simulator) UpdateAgents(){}
func (sim *Simulator) WaitNeighborsSync(){}

func (sim *Simulator) Step(){
	fmt.Printf("Step\n")
	// 隣接エリアとのClock同期をとる

	// 隣接エリアのエージェントを取得

	// エージェント計算を行う

	// 隣接エリアとの同期をとる

	// 計算後のエージェントを取得する

	// Clockを進める
}

func (sim *Simulator) Run() {
	fmt.Printf("Run Simulator\n")
	for i := 0; i < 100; i++ {
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