package ssh_config

import (
	"fmt"
	"log"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

// 1. Root Schema: Represents the entire file
type Config struct {
	Hosts []Host `hcl:"host,block"`
}

// 2. Host Block Schema
type Host struct {
	Name     string `hcl:"name,label"` // Captures "my-service"
	Hostname string `hcl:"hostname"`
	Alias    string `hcl:"alias"`

	// IMPORTANT: This is an Attribute (has '=' in HCL), not a Block.
	// We map it to a struct to decode the object inside.
	Config HostDetails `hcl:"config"`
}

// 3. Config Object Schema
// Because this is decoding an attribute value (Object), we use 'cty' tags
// to map the keys inside the object { ... } to struct fields.
type HostDetails struct {
	User         string `cty:"user"`
	IdentityFile string `cty:"identity_file"`
	Port         int    `cty:"port"`
}

func Parse(input string) (Config, error) {
	var cfg Config
	err := hclsimple.DecodeFile(input, nil, &cfg)
	if err != nil {
		log.Fatalf("Failed to decode: %s", err)
		return Config{}, err
	}
	return cfg, nil
}

func List(input string) {
	cfg, err := Parse(input)
	if err != nil {
		log.Fatalf("Error parsing config: %s", err)
		return
	}

	for _, h := range cfg.Hosts {
		fmt.Printf("Loaded Host: %s\n", h.Name)
		fmt.Printf(" - Hostname: %s\n", h.Hostname)
		fmt.Printf(" - Port:     %d\n", h.Config.Port)
	}
}

func Parser(input string, output string, dryRun bool) (string, error) {

	cfg, err := Parse(input)
	if err != nil {
		return "", err
	}

	for _, h := range cfg.Hosts {
		fmt.Printf("Loaded Host: %s\n", h.Name)
		fmt.Printf(" - Hostname: %s\n", h.Hostname)
		fmt.Printf(" - Port:     %d\n", h.Config.Port)
	}

	return "", nil
}
