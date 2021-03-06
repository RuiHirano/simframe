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

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/otiai10/copy"
)

var (
    builderDirPath string
    currentDirPath string
)

type SimframeConfig struct{
	Name string `json:"name"`
	Version string `json:"version"`
}

func init(){
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        panic("No caller information")
    }
    builderDirPath = path.Dir(path.Dir(filename))+"/builder"
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
	SimframeConfig *SimframeConfig
}

func NewBuilder(id string) *Builder {

	builder := &Builder{
		ID: id,
	}

	builder.SetConfig()
	return builder
}

func (bd *Builder) InstallDockerFile(){
    err := copy.Copy(fmt.Sprintf("%s/Dockerfile", builderDirPath), fmt.Sprintf("%s/build/Dockerfile", currentDirPath))
    if err != nil {
        color.Red("Install dockerfile error %v\n", err,fmt.Sprintf("%s/Dockerfile", builderDirPath))
        os.Exit(1)
    }
}

func (bd *Builder) BuildDockerImage(){
	color.Green("Building docker image...\n")
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	//cmd := exec.Command(builderDirPath+"/docker_build.sh", "sample", "1.0.0", currentDirPath)
	cmd := exec.Command("docker", "image", "build", "-t", fmt.Sprintf("simframe/%s:%s", bd.SimframeConfig.Name, bd.SimframeConfig.Version), "-f", "./build/Dockerfile", ".")
	stdout, err := cmd.StdoutPipe()
	color.Green("Command: %v\n", cmd.String())

	if err != nil {
		color.Red("Error: %v\n", err)
		os.Exit(1)
	}

	cmd.Start()
	s.Start() 

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
	gen := NewGenerator("id", bd.SimframeConfig)
	gen.GenerateResource()
	s.Stop()
}

func (bd *Builder) CreateBuildDirectory(){
	dir := fmt.Sprintf("%s/build", currentDirPath)
	if f, err := os.Stat(dir); os.IsNotExist(err) || !f.IsDir() {
		color.Green("Creating build directory...\n")
		if err := os.Mkdir(dir, 0777); err != nil {
			color.Red("Error: %v\n", err)
			os.Exit(1)
		}
	} 	
}

func (bd *Builder) SetConfig(){
	raw, err := ioutil.ReadFile(fmt.Sprintf("%s/simframe.config.json", currentDirPath))
	if err != nil{
		color.Red("cannot find simframe.config.json\n")
		os.Exit(1)
	}
	sc := &SimframeConfig{}
	err = json.Unmarshal(raw, sc)
	if sc == nil || sc.Name == "" || sc.Version == ""{
		color.Red("invalid args in simframe.config.json\n")
		color.Red("%v", raw, fmt.Sprintf("%s/simframe.config.json", currentDirPath), err)
		os.Exit(1)
	}
	bd.SimframeConfig = sc
}

func (bd *Builder) Build(){
	bd.CreateBuildDirectory()
	bd.InstallDockerFile()
	bd.BuildDockerImage()
	bd.GenerateK8sResource()

	color.Green("Built project!\n\n")
    color.Green("You can run project by below command\n\n")
    color.Green("$ simcli run\n")
}