package ipify

import (
  "github.com/rs/zerolog/log"
  "io"
  "net"
  "net/http"
)

type IP net.IP

func (t IP) GetIP() (net.IP, error) {
  log.Info().Msg("Finding your current IP")
  res, err := http.Get("https://api.ipify.org")
  if err != nil {
    return nil, err
  }
  
  body, err := io.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }
  log.Info().Msgf("Your current IP is: '%s'", string(body))

  return net.ParseIP(string(body)), nil
}