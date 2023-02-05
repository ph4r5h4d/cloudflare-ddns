package cldflr

import (
  "fmt"
  "github.com/cloudflare/cloudflare-go"
  "github.com/ph4r5h4d/cloudflare-ddns/model"
  "github.com/ph4r5h4d/cloudflare-ddns/pkg/dns"
  "net"
)

type DNS struct {
  API *cloudflare.API
  Config model.Cloudflare
}

func (d DNS) Build(config model.Config) (dns.Provider, error) {

  api, err := cloudflare.NewWithAPIToken(config.DnsProviders.Cloudflare.ApiToken)
  if err != nil {
    return DNS{}, err
  }
  d.API = api
  d.Config = config.DnsProviders.Cloudflare
  return d, nil
}

func (d DNS) Update(ip net.IP, records []string) error {
//  ctx := context.Background()
  zone, err := d.API.ZoneIDByName(d.Config.Zone)
  if err != nil {
    return err
  }
  fmt.Println(zone)
  return nil
}