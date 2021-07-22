package api

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewMaster() *Master {

	master := &Master{
		Id: "0",
		ServerAddress: "",
		RegisteredWorkers: []*Worker{},
	}

	return master
}

func (master *Master) Run() {
	fmt.Printf("Run Master\n")
}

func (master *Master) Serve() {
	fmt.Printf("Serve Master\n")

	server := grpc.NewServer()
	svc := NewMasterService(master)
	// 実行したい実処理をseverに登録する
	RegisterMasterServiceServer(server, svc)

	reflection.Register(server)
	lis, err := net.Listen("tcp", ":9000")
	defer lis.Close()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Serving! port is 9000\n")
	server.Serve(lis)
}

func (master *Master) RegisterWorker(wk *Worker){
	master.RegisteredWorkers = append(master.RegisteredWorkers, wk)
	fmt.Print(master.RegisteredWorkers)
}

func (master *Master) UpdateWorker(wk *Worker){
	for i, v := range master.RegisteredWorkers{
		if v.Id == wk.Id{
			master.RegisteredWorkers[i] = wk
			fmt.Print("update")
			return
		}
	}
	fmt.Print("Error: worker is not registerd...")
}

type MasterService struct{
 	master *Master
	UnimplementedMasterServiceServer
}

func NewMasterService(master *Master) *MasterService{
	ms := &MasterService{
		master: master,
	}
	return ms
}

func (b *MasterService) RegisterWorker(ctx context.Context, request *RegisterWorkerRequest) (*RegisterWorkerResponse, error) {
	fmt.Printf("getRequest %v\n", request)

	wk := NewWorker()
	b.master.RegisterWorker(wk)
	response := &RegisterWorkerResponse{}
	return response, nil
}

func (b *MasterService) UpdateWorker(ctx context.Context, request *UpdateWorkerRequest) (*UpdateWorkerResponse, error) {
	fmt.Printf("getRequest %v\n", request)

	wk := NewWorker()
	b.master.UpdateWorker(wk)
	response := &UpdateWorkerResponse{}
	return response, nil
}
