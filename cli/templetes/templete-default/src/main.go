package main

import (
	"github.com/RuiHirano/simframe/app"
	"github.com/RuiHirano/simframe/engine"
)

// TODO: Hide here from user
func main() {
	ap := App{app.NewApp()}

	en := engine.NewEngine(ap)
	en.Run()
}