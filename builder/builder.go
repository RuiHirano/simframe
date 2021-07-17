package builder

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"time"

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
	cmd := exec.Command(builderDirPath+"/docker_build.sh", "sample", "1.0.0", builderDirPath)
	stdout, err := cmd.StdoutPipe()

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

func (bd *Builder) Build(){
	bd.CreateBuildDirectory()
	bd.BuildDockerImage()
	bd.GenerateK8sResource()
}