package main

import (
	"github.com/RuiHirano/simframe/config"
	"github.com/RuiHirano/simframe/engine"
	"github.com/RuiHirano/simframe/scenario"
	myscenario "github.com/RuiHirano/simframe/src/scenario"
)

// TODO: Hide here from user
func main() {

	// docker内でここを実行する。

	sn := myscenario.NewScenario1()

	config := config.NewConfig(&config.ConfigParams{
		ServerNum: 1,
	})

	en := engine.NewEngine([]scenario.IScenario{sn}, config)
	en.Run()
}