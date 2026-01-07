package cmd

import (
	"fmt"
	"os"

	"github.com/gjsstigter/ssh_config/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ssh_config",
	Short: "A CLI tool to manage ssh config files",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Initialize()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
