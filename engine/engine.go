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

type IEngine interface {
	Run(runType string)
}

type Engine struct {
	App app.IApp
}

func NewEngine(ap app.IApp) *Engine {

	engine := &Engine{
		App: ap,
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
	case "SIMULATOR":
		sim := NewSimulator()
		sim.Serve()
		sim.Run()	
	}
}


func (en *Engine) Serve() {
	fmt.Printf("Serve Engine\n")

	server := grpc.NewServer()
	svc := NewEngineService(en.RegisterSimulatorHandler)
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


func (engine *Engine) RegisterSimulatorHandler() (app.IArea, app.IClock, []app.IAgent){
	fmt.Printf("Register simulator\n")
	return app.NewArea(), app.NewClock(), []app.IAgent{}
}
