package engine

import (
	"fmt"
	"log"
	"net"

	"github.com/RuiHirano/simframe/api"

	"github.com/RuiHirano/simframe/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
}

func (sim *Simulator) Serve() {
	fmt.Printf("Serve Simulator\n")

	server := grpc.NewServer()
	svc := NewSimulatorService(sim.RunSimulatorHandler, sim.GetNeighborAgentsHandler)
	// 実行したい実処理をseverに登録する
	api.RegisterSimulatorServiceServer(server, svc)

	reflection.Register(server)
	lis, err := net.Listen("tcp", ":9000")
	defer lis.Close()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Serving! port is 9000\n")
	server.Serve(lis)
}


func (sim *Simulator) RunSimulatorHandler(){
	fmt.Printf("Run simulator\n")
	sim.Run()
}

func (sim *Simulator) GetNeighborAgentsHandler(){
	fmt.Printf("Get Neighbor Agents\n")
	
}