package ssh_config

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func generateHostnames(hostname string) []string {
	var hostnames []string
	patterns := []struct {
		re  *regexp.Regexp
		typ string
	}{
		{regexp.MustCompile(`\[(\d+)\.\.(\d+)\]`), "range"},
		{regexp.MustCompile(`\[([a-zA-Z0-9-|]+(?:\.[a-zA-Z0-9-|]+)*)+\]`), "variation"},
		{regexp.MustCompile(`^([a-zA-Z0-9_]|[a-zA-Z0-9_][a-zA-Z0-9-_]{0,61}[a-zA-Z0-9_])` +
			`(\.([a-zA-Z0-9_]|[a-zA-Z0-9_][a-zA-Z0-9-_]{0,61}[a-zA-Z0-9_]))*$`), "hostname"},
	}

	for _, p := range patterns {
		if subs := p.re.FindStringSubmatch(hostname); subs != nil {
			switch p.typ {
			case "range":
				start, _ := strconv.Atoi(subs[1])
				end, _ := strconv.Atoi(subs[2])
				for i := start; i <= end; i++ {
					hostnames = append(hostnames, p.re.ReplaceAllString(hostname, strconv.Itoa(i)))
				}
				return hostnames
			case "list", "brace":
				parts := strings.Split(subs[1], ",")
				for _, part := range parts {
					hostnames = append(hostnames, p.re.ReplaceAllString(hostname, part))
				}
				return hostnames
			}
		}
	}
	return []string{hostname}
}

func Translate(host Host) ([]SSHConfigHost, error) {
	sshConfigHosts := []SSHConfigHost{}
	if host.Hostname != "" {
		hostnames := generateHostnames(host.Hostname)
		sshConfigHost := SSHConfigHost{
			Hosts:    hostnames,
			Hostname: host.Name,
			Port:     host.Config.Port,
			User:     host.Config.User,
		}
		sshConfigHosts = append(sshConfigHosts, sshConfigHost)
		return sshConfigHosts, nil
	}
	return []SSHConfigHost{}, fmt.Errorf("Hostname is empty for host: %s", host.Name)
}
