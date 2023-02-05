package main

import (
	"errors"
	"github.com/ph4r5h4d/cloudflare-ddns/model"
	"github.com/ph4r5h4d/cloudflare-ddns/pkg/dns"
	cldflr "github.com/ph4r5h4d/cloudflare-ddns/pkg/dns/cloudflare"
	ip "github.com/ph4r5h4d/cloudflare-ddns/pkg/ip/current"
	"github.com/ph4r5h4d/cloudflare-ddns/pkg/ip/current/icanhazip"
	"github.com/ph4r5h4d/cloudflare-ddns/pkg/ip/current/ipify"

	"github.com/spf13/viper"
)

type Dependencies struct {
	Config      model.Config
	IPInterface ip.GetIP
	DNSProvider dns.Provider
}

var ipProviders = map[string]interface{}{
	"ipify":     &ipify.IP{},
	"icanhazip": &icanhazip.IP{},
}

var dnsProvider = map[string]interface{}{
	"cloudflare": cldflr.DNS{},
}

func setupConfiguration() (model.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return model.Config{}, err
	}

	c := &model.Config{}
	err = viper.Unmarshal(c)
	if err != nil {
		return model.Config{}, err
	}
	return *c, nil
}

func setupGetIPProvider(config model.Config) (ip.GetIP, error) {
	ipp, ok := ipProviders[config.IPProvider]
	if !ok {
		return nil, errors.New("invald IP provider")
	}
	return ipp.(ip.GetIP), nil
}

func setupDNSProvider(config model.Config) (dns.Provider, error) {
	dnsp, ok := dnsProvider[config.DnsProvider]
	if !ok {
		return nil, errors.New("invalid DNS provider")
	}
	return dnsp.(dns.Provider), nil
}

func setupDependencies() (Dependencies, error) {
	c, err := setupConfiguration()
	if err != nil {
		return Dependencies{}, nil
	}

	ip, err := setupGetIPProvider(c)
	if err != nil {
		return Dependencies{}, nil
	}

	dnsImpl, err := setupDNSProvider(c)
	if err != nil {
		return Dependencies{}, nil
	}
	dnsp, err := dnsImpl.Build(c)
	if err != nil {
		return Dependencies{}, nil
	}

	d := Dependencies{
		Config:      c,
		IPInterface: ip,
		DNSProvider: dnsp,
	}
	return d, err
}
