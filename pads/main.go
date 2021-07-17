package main

import (
	"github.com/RuiHirano/simframe/config"
	"github.com/RuiHirano/simframe/engine/master"
	"github.com/RuiHirano/simframe/engine/worker"
	"github.com/RuiHirano/simframe/scenario"
)

type IEngine interface {
	Run()
}

type Engine struct {
	ID string 
	Scenarios []scenario.IScenario
	Config config.IConfig
}

func NewEngine(scenarios []scenario.IScenario, conf config.IConfig) *Engine {

	engine := &Engine{
		ID: "0",
		Scenarios: scenarios,
		Config: conf,
	}

	return engine
}

func (engine *Engine) Run() {
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

func main(){
	sns := []scenario.IScenario{}
	conf := config.NewConfig(&config.ConfigParams{ServerNum: 1})
	engine := NewEngine(sns, conf)
	engine.Run()
}