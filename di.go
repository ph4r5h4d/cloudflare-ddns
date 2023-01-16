package main

import (
	"errors"
  "github.com/ph4r5h4d/cloudflare-ddns/model"
	"github.com/ph4r5h4d/cloudflare-ddns/pkg/ip/current"
  "github.com/ph4r5h4d/cloudflare-ddns/pkg/ip/current/icanhazip"
  "github.com/ph4r5h4d/cloudflare-ddns/pkg/ip/current/ipify"
  "github.com/spf13/viper"
)

var ipProviders = map[string]interface{}{
  "ipify": &ipify.IP{},
  "icanhazip": &icanhazip.IP{},
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

func setupGetIPProvider(config model.Config) (current.GetIP, error) {
  s, ok := ipProviders[config.IPProvider]
	if !ok {
		return nil, errors.New("invald IP provider")
	}
	return s.(current.GetIP), nil
}

func setupDependencies() (model.Dependencies, error) {
	c, err := setupConfiguration()
  if err != nil {
    return model.Dependencies{}, nil
  }
	ip, err := setupGetIPProvider(c)
  if err != nil {
    return model.Dependencies{}, nil
  }
	d := model.Dependencies{
		Config:      c,
		IPInterface: ip,
	}
	return d, err
}
