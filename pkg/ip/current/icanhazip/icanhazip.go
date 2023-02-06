package icanhazip

import (
  "github.com/rs/zerolog/log"
  "io"
  "net"
  "net/http"
  "strings"
)

type IP net.IP

func (t IP) GetIP() (net.IP, error) {
  log.Info().Msg("Finding your current IP")
  res, err := http.Get("https://icanhazip.com")
  if err != nil {
    return nil, err
  }
  body, err := io.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }
  currentIp := net.ParseIP(strings.Replace(string(body), "\n", "", -1))
  log.Info().Msgf("Your current IP is: '%s'", currentIp)
  return currentIp, nil
}