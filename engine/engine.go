package engine

import (
	"fmt"

	ap "github.com/RuiHirano/simframe/app"
	"github.com/RuiHirano/simframe/builder"
	"github.com/RuiHirano/simframe/engine/pads/master"
	"github.com/RuiHirano/simframe/engine/pads/worker"
)

type IEngine interface {
	Build()
	RunPads()
}

type Engine struct {
	App ap.App
}

func NewEngine(app ap.App) *Engine {

	engine := &Engine{
		App: app,
	}

	return engine
}

func (engine *Engine) RunPads() {
	runType := "WORKER"
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

func (engine *Engine) Build() {
	fmt.Printf("build")
	bd := builder.Builder
	bd.Build(engine.App)
}


func main(){
	fmt.Printf("main")
	
}