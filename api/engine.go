package api

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewEngine(app *App) *Engine {

	engine := &Engine{
		App: app,
	}

	return engine
}

func (engine *Engine) Run(runType string) {
	//unType := "WORKER"
	switch runType {
	case "ENGINE":
		engine.Serve()
		// calc use resources

		// create resources

		// apply resources
	case "MASTER":
		master := NewMaster()
		master.Serve()
		master.Run()
	case "WORKER":
		worker := NewWorker()
		worker.Serve()
		worker.Run()	
	}
}


func (engine *Engine) Serve() {
	fmt.Printf("Serve Engine\n")

	server := grpc.NewServer()
	svc := NewEngineService(engine)
	// 実行したい実処理をseverに登録する
	RegisterEngineServiceServer(server, svc)

	reflection.Register(server)
	lis, err := net.Listen("tcp", ":10000")
	defer lis.Close()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Serving! port is 10000\n")
	server.Serve(lis)
}


type EngineService struct{
	engine *Engine
	UnimplementedEngineServiceServer
}

func NewEngineService(engine *Engine) *EngineService{
   ms := &EngineService{
	   engine: engine,
   }
   return ms
}

func (es *EngineService) GetStatus(ctx context.Context, request *GetStatusRequest) (*GetStatusResponse, error) {
   fmt.Printf("getRequest %v\n", request)

   response := &GetStatusResponse{
		Config: &Config{},
		Scenarios: []*Scenario{},
   }
   return response, nil
}
