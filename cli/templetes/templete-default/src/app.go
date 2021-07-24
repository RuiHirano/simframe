package myapp

import (
	myscenario "scenario"

	"github.com/RuiHirano/simframe/app"
)

func Scenarios() []app.IScenario{
	sn := myscenario.NewScenario1()
	return []app.IScenario{sn}
}
