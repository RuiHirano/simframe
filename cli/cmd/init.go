package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
    Use:   "init",
    Short: "command line calculator",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Init command")
    },
}

func init() {
	RootCmd.AddCommand(initCmd)
}