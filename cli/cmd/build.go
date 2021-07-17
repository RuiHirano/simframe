package cmd

import (
	"fmt"

	"github.com/RuiHirano/simframe/builder"
	"github.com/spf13/cobra"
)

func Build(){
    bd := builder.NewBuilder("id")
    bd.Build()
}

var buildCmd = &cobra.Command{
    Use:   "build",
    Short: "command line calculator",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Build command")
        Build()
    },
}

func init() {
	RootCmd.AddCommand(buildCmd)
}