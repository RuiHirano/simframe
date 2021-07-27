package main


import (
	myapp "app"
	"flag"
	"os"
	"strconv"

	"github.com/fatih/color"

	"github.com/RuiHirano/simframe/app"
	"github.com/RuiHirano/simframe/engine"
)


// TODO: Hide here from user
func main() {
	

	sns :=  myapp.Scenarios()
	conf :=  app.NewConfig()
	ap := app.NewApp(sns, conf)

	en := engine.NewEngine(ap)

	flag.Parse()
    arg := flag.Arg(0)
    runType := flag.Arg(1)
    id := flag.Arg(2)
    port, _ := strconv.Atoi(flag.Arg(3))
	switch arg {
	case "run":	
		switch runType {
		case "engine":	
			en.Run("ENGINE", "engine", 10000)
		case "simulator":	
			en.Run("SIMULATOR", id, port)
		}
	}
	color.Red("invalid arg in engine\n")
	os.Exit(1)
}