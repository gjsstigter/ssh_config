package ssh_config

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

type SSHConfig struct {
	Hosts []SSHConfigHost
}

type SSHConfigHost struct {
	Hosts    []string
	Hostname string
	Port     int
	User     string
}
