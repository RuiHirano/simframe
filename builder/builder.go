package builder

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"runtime"
	"time"

	"github.com/RuiHirano/simframe/util"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

var (
    builderDirPath string
    currentDirPath string
)

func init(){
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        panic("No caller information")
    }
    builderDirPath = path.Dir(path.Dir(filename) + "/../builder/")
    p, _ := os.Getwd()
    currentDirPath = p
}

type IBuilder interface {
	Build()
	Scenarios()
	Config()
}

type Builder struct {
	ID string
}

func NewBuilder(id string) *Builder {

	builder := &Builder{
		ID: id,
	}

	return builder
}

func (bd *Builder) BuildDockerImage(){
	color.Green("Building docker image...\n")
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	//cmd := exec.Command(builderDirPath+"/docker_build.sh", "sample", "1.0.0", currentDirPath)
	cmd := exec.Command(fmt.Sprintf("docker build -t simframe/%s:%s %s", "sample", "1.0.0", currentDirPath))
	stdout, err := cmd.StdoutPipe()
	color.Green("Command: %v\n", cmd.String())

	if err != nil {
		color.Red("Error: %v\n", err)
		os.Exit(1)
	}

	s.Start() 
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Println()
	}

	cmd.Wait()
	s.Stop()
}

func (bd *Builder) GenerateK8sResource(){
	color.Green("Generating kubernetes resource file...\n")
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start() 
	gen := NewGenerator("id")
	gen.GenerateResource()
	s.Stop()
}

func (bd *Builder) CreateBuildDirectory(){
	if err := os.Mkdir(fmt.Sprintf("%s/build", currentDirPath), 0777); err != nil {
        color.Red("Error: %v\n", err)
		os.Exit(1)
    }
}

func (bd *Builder) GetApp(){
	raw, err := ioutil.ReadFile(fmt.Sprintf("%s/simframe.config.json", currentDirPath))
	if err != nil{
		color.Red("cannot find simframe.config.json\n")
		os.Exit(1)
	}
	var sc *util.SimFrameConfig
	json.Unmarshal(raw, sc)

	//entry := sc.Entry
	//fmt.Printf(entry)
}


func (bd *Builder) Build(){
	bd.GetApp()
	bd.CreateBuildDirectory()
	bd.BuildDockerImage()
	bd.GenerateK8sResource()

	color.Green("Built project!\n\n")
    color.Green("You can run project by below command\n\n")
    color.Green("$ simcli run\n")
}