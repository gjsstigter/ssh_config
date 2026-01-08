package ssh_config

import (
	"log"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

func Parse(input string) (Config, error) {
	var cfg Config
	err := hclsimple.DecodeFile(input, nil, &cfg)
	if err != nil {
		log.Fatalf("Failed to decode: %s", err)
		return Config{}, err
	}
	return cfg, nil
}
