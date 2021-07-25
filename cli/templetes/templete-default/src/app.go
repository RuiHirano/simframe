package myapp

import (
	myscenario "scenario"

	"github.com/RuiHirano/simframe/app"
)

func Scenarios() []app.IScenario{
	sn := myscenario.NewScenario1()

	for _, ag := range sn.GetAgents() {
		ag.Step()
	}
	return []app.IScenario{sn}
}
