package engine

import (
	"fmt"

	ap "github.com/RuiHirano/simframe/app"
	"github.com/RuiHirano/simframe/engine/master"
	"github.com/RuiHirano/simframe/engine/worker"
)

type IEngine interface {
	RunPads()
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

func (engine *Engine) RunPads(runType string) {
	//unType := "WORKER"
	if runType == "MASTER"{
		master := master.NewMaster()
		master.Serve()
		master.Run()
	}else if runType == "WORKER"{
		worker := worker.NewWorker()
		worker.Serve()
		worker.Run()
	}
}

func main(){
	fmt.Printf("main")
	
}