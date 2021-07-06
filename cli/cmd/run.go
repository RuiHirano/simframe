package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
    Use:   "run",
    Short: "command line calculator",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("run command")
    },
}

func init() {
	RootCmd.AddCommand(runCmd)
}