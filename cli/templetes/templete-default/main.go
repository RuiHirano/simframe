package main


import (
	myapp "app"
	"flag"
	"os"

	"github.com/fatih/color"

	"github.com/RuiHirano/simframe/app"
	"github.com/RuiHirano/simframe/engine"
)



// TODO: Hide here from user
func main() {
	

	sns :=  myapp.Scenarios()
	config := myapp.Config()
	ap := app.NewApp(sns, config)

	en := engine.NewEngine(ap)

	flag.Parse()
    arg := flag.Arg(0)
    runType := flag.Arg(1)
	switch arg {
	case "run":	
		switch runType {
		case "master":	
			en.RunPads("MASTER")
		case "worker":	
			en.RunPads("WORKER")
		}
	}
	color.Red("invalid arg in engine\n")
	os.Exit(1)
}