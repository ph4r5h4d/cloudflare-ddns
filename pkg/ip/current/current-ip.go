package current

import (
  "net"
)
type GetIP interface {
  GetIP() (net.IP, error)
}