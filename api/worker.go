package api

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func NewWorker() *Worker {

	worker := &Worker{
		Id: "0",
		ServerAddress: "",
		NeighborWorkers: []*Worker{},
	}

	return worker
}

func (worker *Worker) GetNeighborData() []*Agent{
	agents := []*Agent{}
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
	RegisterWorkerServiceServer(server, svc)

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
	worker *Worker
	UnimplementedWorkerServiceServer
}

func NewWorkerService(worker *Worker) *WorkerService{
	ws := &WorkerService{
		worker: worker,
	}
	return ws
}

func (b *WorkerService) RunWorker(ctx context.Context, request *RunWorkerRequest) (*RunWorkerResponse, error) {
	fmt.Printf("getRequest %v\n", request)

	response := &RunWorkerResponse{}
	return response, nil
}

func (b *WorkerService) GetAgents(ctx context.Context, request *GetAgentsRequest) (*GetAgentsResponse, error) {
	fmt.Printf("getRequest %v\n", request)

	response := &GetAgentsResponse{}
	return response, nil
}
