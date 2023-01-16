package icanhazip

import (
  "io"
  "net"
  "net/http"
  "strings"
)

type IP net.IP

func (t IP) GetIP() (net.IP, error) {
  res, err := http.Get("https://icanhazip.com")
  if err != nil {
    return nil, err
  }
  body, err := io.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }
  return net.ParseIP(strings.Replace(string(body), "\n", "", -1)), nil
}