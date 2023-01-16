package model

import ip "github.com/ph4r5h4d/cloudflare-ddns/pkg/ip/current"

type Dependencies struct {
  Config
  IPInterface ip.GetIP
}