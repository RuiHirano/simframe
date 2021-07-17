package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Get(){
    fmt.Println("Get command")
}

var getCmd = &cobra.Command{
    Use:   "get",
    Short: "command line getter",
    Run: func(cmd *cobra.Command, args []string) {
        Get()
    },
}

func GetScenario(){
    fmt.Println("Get Scenario command")
}

var scenarioCmd = &cobra.Command{
    Use:   "scenario",
    Short: "command line getter",
    Run: func(cmd *cobra.Command, args []string) {
        GetScenario()
    },
}

func init() {
    getCmd.AddCommand(scenarioCmd)
	RootCmd.AddCommand(getCmd)
}