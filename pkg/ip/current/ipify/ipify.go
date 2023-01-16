package ipify

import (
  "io"
  "net"
  "net/http"
)

type IP net.IP

func (t IP) GetIP() (net.IP, error) {
  res, err := http.Get("https://api.ipify.org")
  if err != nil {
    return nil, err
  }
  body, err := io.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }
  return net.ParseIP(string(body)), nil
}