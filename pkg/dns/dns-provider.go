package dns

import (
  "github.com/ph4r5h4d/cloudflare-ddns/model"
  "net"
)

type Provider interface {
  Build(config model.Config) (Provider, error)
  Update(ip net.IP, records []string) error
}