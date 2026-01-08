package ssh_config

import (
	"log"
)

func List(input string) {
	cfg, err := Parse(input)
	if err != nil {
		log.Fatalf("Error parsing config: %s", err)
		return
	}

	for _, h := range cfg.Hosts {
		sshConfigHosts, err := Translate(h)
		if err != nil {
			log.Printf("Error translating host %s: %s", h.Name, err)
			continue
		}
		for _, sch := range sshConfigHosts {
			log.Printf("Host: %s, Hostname: %s, Port: %d", h.Name, sch.Hostname, sch.Port)
		}
	}
}
