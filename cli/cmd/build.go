package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
    Use:   "build",
    Short: "command line calculator",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Build command")
    },
}

func init() {
	RootCmd.AddCommand(buildCmd)
}