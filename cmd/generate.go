package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate SSH config files",
	Long:  "Generate SSH config files from templates and user input",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating SSH config files...")

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
