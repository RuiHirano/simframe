package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    Use:   "simcli",
    Short: "command line calculator",
}
