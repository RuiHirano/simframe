package worker

import (
	"context"
	"fmt"
	"log"
	"net"

	api "github.com/RuiHirano/simframe/pads/proto"

	"github.com/RuiHirano/simframe/app/model"
	"github.com/RuiHirano/simframe/app/scenario"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type IWorker interface {
	GetNeighborData() []model.IAgent
	StepAgents()
	UpdateAgents()
	ClockForward()
	WaitNeighborsSync()
	Step()
	Serve()
	GetID() string
	Run()
}

type Worker struct {
	ID string 
	Scenario scenario.IScenario
}

func NewWorker() *Worker {
	agents := []model.IAgent{}
	area := scenario.NewArea()
	clock := scenario.NewClock()

	sn := scenario.NewScenario(agents, area, clock)

	worker := &Worker{
		ID: "0",
		Scenario: sn,
	}

	return worker
}

func (worker *Worker) GetID() string{
	return worker.ID
}

func (worker *Worker) GetNeighborData() []model.IAgent{
	agents := []model.IAgent{}
	return agents
}

func (worker *Worker) StepAgents(){}
func (worker *Worker) ClockForward(){}
func (worker *Worker) UpdateAgents(){}
func (worker *Worker) WaitNeighborsSync(){}

func (worker *Worker) Step(){
	// 隣接エリアとのClock同期をとる

	// 隣接エリアのエージェントを取得

	// エージェント計算を行う

	// 隣接エリアとの同期をとる

	// 計算後のエージェントを取得する

	// Clockを進める
}

func (worker *Worker) Run() {
	fmt.Printf("Run Worker\n")
	for i := 0; i < 100; i++ {
		worker.Step()
	}
}

func (worker *Worker) Serve() {
	fmt.Printf("Serve Worker\n")

	server := grpc.NewServer()
	svc := NewWorkerService(worker)
	// 実行したい実処理をseverに登録する
	api.RegisterWorkerServiceServer(server, svc)

	reflection.Register(server)
	lis, err := net.Listen("tcp", ":9000")
	defer lis.Close()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Serving! port is 9000\n")
	server.Serve(lis)
}


type WorkerService struct{
	worker IWorker
}

func NewWorkerService(worker IWorker) *WorkerService{
	ws := &WorkerService{
		worker: worker,
	}
	return ws
}

func (b *WorkerService) RunWorker(ctx context.Context, request *api.RunWorkerRequest) (*api.RunWorkerResponse, error) {
	fmt.Printf("getRequest %v\n", request)

	response := &api.RunWorkerResponse{}
	return response, nil
}

func (b *WorkerService) GetAgents(ctx context.Context, request *api.GetAgentsRequest) (*api.GetAgentsResponse, error) {
	fmt.Printf("getRequest %v\n", request)

	response := &api.GetAgentsResponse{}
	return response, nil
}
