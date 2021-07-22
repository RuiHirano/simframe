package engine

import (
	"context"
	"fmt"
	"log"
	"net"

	ap "github.com/RuiHirano/simframe/app"
	"github.com/RuiHirano/simframe/engine/api"
	"github.com/RuiHirano/simframe/engine/master"
	"github.com/RuiHirano/simframe/engine/worker"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type IEngine interface {
	Run(runType string)
}

type Engine struct {
	App ap.IApp
}

func NewEngine(app ap.IApp) *Engine {

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
		master := master.NewMaster()
		master.Serve()
		master.Run()
	case "WORKER":
		worker := worker.NewWorker()
		worker.Serve()
		worker.Run()	
	}
}


func (engine *Engine) Serve() {
	fmt.Printf("Serve Engine\n")

	server := grpc.NewServer()
	svc := NewEngineService(engine)
	// 実行したい実処理をseverに登録する
	api.RegisterEngineServiceServer(server, svc)

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
	engine IEngine
	api.UnimplementedEngineServiceServer
}

func NewEngineService(engine IEngine) *EngineService{
   ms := &EngineService{
	   engine: engine,
   }
   return ms
}

func (es *EngineService) GetStatus(ctx context.Context, request *api.GetStatusRequest) (*api.GetStatusResponse, error) {
   fmt.Printf("getRequest %v\n", request)

   response := &api.GetStatusResponse{
	   Status: &api.SimulationStatus{
		Config: &api.Config{},
		Scenarios: []*api.Scenario{},
	   },
   }
   return response, nil
}
