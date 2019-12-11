package node

import "net"

type Node struct {
	Instance string
	Addr     net.IP
	Port     int
}
