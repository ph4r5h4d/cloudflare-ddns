package model

// Config struct holds all the application configuration
type Config struct {
  DnsProvider  string       `yaml:"dnsProvider"`
  DnsProviders DnsProviders `yaml:"dnsProviders"`
  IPProvider   string       `yaml:"IPProvider"`
  Records      []string     `yaml:"records"`
}

// DnsProviders list of all supported providers
type DnsProviders struct {
  Cloudflare Cloudflare `yaml:"cloudflare"`
}

// Cloudflare configuration
type Cloudflare struct {
  ApiKey   string `yaml:"apiKey"`
  ApiEmail string `yaml:"apiEmail"`
}

