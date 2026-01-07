package cmd

import (
	"fmt"

	"github.com/gjsstigter/ssh_config/src/ssh_config"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List SSH config files",
	Long:  "List SSH config files from templates and user input",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating SSH config files...")
		input, _ := cmd.Flags().GetString("input")

		fmt.Printf("Input: %s\n", input)

		// Add logic to generate SSH config files here
		ssh_config.List(input)
	},
}

func init() {
	listCmd.Flags().StringP("input", "i", "~/.ssh_config", "Input folder")
	rootCmd.AddCommand(listCmd)
}
