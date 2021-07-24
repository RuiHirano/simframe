package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func DeleteResources(){
    color.Green("Deleting kubernetes resources...\n")
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	cmd := exec.Command("kubectl", "delete", "pods,services", "-l", "app=simframe")
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

func Delete(){
    DeleteResources()
}

var deleteCmd = &cobra.Command{
    Use:   "delete",
    Short: "command line calculator",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("delete command")
        Delete()
    },
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}