package cmd

import (
	"fmt"

	"github.com/gjsstigter/ssh_config/src/ssh_config"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate SSH config files",
	Long:  "Generate SSH config files from templates and user input",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating SSH config files...")
		input, _ := cmd.Flags().GetString("input")
		output, _ := cmd.Flags().GetString("output")
		dryRun, _ := cmd.Flags().GetBool("dry-run")

		fmt.Printf("Input: %s\n", input)
		fmt.Printf("Output: %s\n", output)
		fmt.Printf("Dry Run: %v\n", dryRun)

		// Add logic to generate SSH config files here
		err := ssh_config.Parser(input, output, dryRun)
		if err != nil {
			fmt.Printf("Error generating SSH config: %v\n", err)
		}
	},
}

func init() {
	generateCmd.Flags().StringP("output", "o", "./output/ssh_config", "Output file for the generated SSH config")
	generateCmd.Flags().StringP("input", "i", "~/.ssh_config", "Input folder")
	generateCmd.Flags().BoolP("dry-run", "d", true, "Dry run")
	rootCmd.AddCommand(generateCmd)
}
