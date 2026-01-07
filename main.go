package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gjsstigter/ssh_config/cmd"
	"github.com/gjsstigter/ssh_config/config"
)

func main() {
	fmt.Println("=== SSH Config Parser ===")
	config.Initialize()

	if err := config.LoadConfig(); err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	log.SetFlags(0)
	cmd.Execute()
}
