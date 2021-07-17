package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
    buildDirPath string
)

func init(){
    p, _ := os.Getwd()
    buildDirPath = path.Dir(p + "/build/")
    fmt.Printf(buildDirPath)
}

func ApplyResources(){
    color.Green("Applying kubernetes resources...\n")
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	cmd := exec.Command("kubectl", "apply", "-f", fmt.Sprintf("%s/app.yaml", buildDirPath))
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

func CheckBuildDir(){
    if _, err := os.Stat(buildDirPath); err != nil {
        color.Red("Error: %v\n", err)
        color.Yellow("Did you build project? -> Please run 'simcli build'\n")
        color.Yellow("Did you run command at project root? -> Please run at project root\n")
		os.Exit(1)
    }
    if _, err := os.Stat(buildDirPath+"/app.yaml"); err != nil {
        color.Red("Error: %v\n", err)
        color.Yellow("Did you build project? -> Please run 'simcli build'\n")
		os.Exit(1)
    }
}

func Run(){
    CheckBuildDir()
    ApplyResources()
}

var runCmd = &cobra.Command{
    Use:   "run",
    Short: "command line calculator",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("run command")
        Run()
    },
}

func init() {
	RootCmd.AddCommand(runCmd)
}