package master

import (
	"context"
	"fmt"
	"log"
	"net"

	api "github.com/RuiHirano/simframe/pads/proto"
	"github.com/RuiHirano/simframe/pads/worker"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type IMaster interface {
	Run()
	Serve()
	RegisterWorker(worker.IWorker)
	UpdateWorker(worker.IWorker)
}

type Master struct {
	ID string 
	Workers []worker.IWorker
}

func NewMaster() *Master {

	master := &Master{
		ID: "0",
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
	api.RegisterMasterServiceServer(server, svc)

	reflection.Register(server)
	lis, err := net.Listen("tcp", ":9000")
	defer lis.Close()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Serving! port is 9000\n")
	server.Serve(lis)
}

func (master *Master) RegisterWorker(wk worker.IWorker){
	master.Workers = append(master.Workers, wk)
	fmt.Print(master.Workers)
}

func (master *Master) UpdateWorker(wk worker.IWorker){
	for i, v := range master.Workers{
		if v.GetID() == wk.GetID(){
			master.Workers[i] = wk
			fmt.Print("update")
			return
		}
	}
	fmt.Print("Error: worker is not registerd...")
}

type MasterService struct{
 	master IMaster
}

func NewMasterService(master IMaster) *MasterService{
	ms := &MasterService{
		master: master,
	}
	return ms
}

func (b *MasterService) RegisterWorker(ctx context.Context, request *api.RegisterWorkerRequest) (*api.RegisterWorkerResponse, error) {
	fmt.Printf("getRequest %v\n", request)

	wk := worker.NewWorker()
	b.master.RegisterWorker(wk)
	response := &api.RegisterWorkerResponse{}
	return response, nil
}

func (b *MasterService) UpdateWorker(ctx context.Context, request *api.UpdateWorkerRequest) (*api.UpdateWorkerResponse, error) {
	fmt.Printf("getRequest %v\n", request)

	wk := worker.NewWorker()
	b.master.UpdateWorker(wk)
	response := &api.UpdateWorkerResponse{}
	return response, nil
}
