package node

import (
	"net"

	"github.com/dmytro-kolesnyk/dds/connection"
)

type Node struct {
	Instance string
	Addr     net.IP
	Port     int
	Conn     connection.Connection
}
