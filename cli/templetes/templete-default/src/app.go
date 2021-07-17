package main

import (
	"github.com/RuiHirano/simframe/builder"
	"github.com/RuiHirano/simframe/config"
	"github.com/RuiHirano/simframe/scenario"
	myscenario "github.com/RuiHirano/simframe/src/scenario"
)

type Builder struct{
	builder.Builder
}

func (bd *Builder) Scenarios() []scenario.IScenario{
	sn := myscenario.NewScenario1()
	return []scenario.IScenario{sn}
}

func (bd *Builder) Config() config.IConfig{
	config := config.NewConfig(&config.ConfigParams{
		ServerNum: 1,
	})
	return config
}