package ssh_config

import (
	"fmt"
	"os"
	"strings"
)

// Parser converts input SSH config to output format
func Parser(inputPath string, outputPath string, dryRun bool) error {
	// Parse the input config file
	cfg, err := Parse(inputPath)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	// Build the output content
	var output strings.Builder

	for _, host := range cfg.Hosts {
		// Translate host configuration
		sshHosts, err := Translate(host)
		if err != nil {
			return fmt.Errorf("failed to translate host: %w", err)
		}

		// Remove duplicate hostnames
		sshHosts = removeDuplicates(sshHosts)

		// Generate SSH config entries
		for _, host := range sshHosts {
			writeHostEntry(&output, host)
		}
	}

	// Print to console
	fmt.Print(output.String())

	// Write to file if not dry run
	if dryRun {
		fmt.Println("Dry run enabled, not writing to file.")
		return nil
	}

	if err := os.WriteFile(outputPath, []byte(output.String()), 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	fmt.Printf("Successfully wrote to: %s\n", outputPath)
	return nil
}

// removeDuplicates filters out SSH hosts with duplicate hostnames
func removeDuplicates(hosts []SSHConfigHost) []SSHConfigHost {
	seen := make(map[string]bool)
	unique := []SSHConfigHost{}

	for _, host := range hosts {
		if !seen[host.Hostname] {
			seen[host.Hostname] = true
			unique = append(unique, host)
		}
	}

	return unique
}

// writeHostEntry formats and writes a single SSH host entry
func writeHostEntry(output *strings.Builder, host SSHConfigHost) {
	output.WriteString(fmt.Sprintf("Host %s\n", strings.Join(host.Hosts, " ")))
	output.WriteString(fmt.Sprintf("  Hostname %s\n", host.Hostname))

	if host.Port != 0 {
		output.WriteString(fmt.Sprintf("  Port %d\n", host.Port))
	}

	if host.User != "" {
		output.WriteString(fmt.Sprintf("  User %s\n", host.User))
	}

	output.WriteString("\n")
}
