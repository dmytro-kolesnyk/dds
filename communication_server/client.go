package communicationServer

import (
	"fmt"
	"sync"

	"github.com/dmytro-kolesnyk/dds/node"
)

type neighbours struct {
	nodes map[string]*node.Node
	//nodes []*node.Node

	mux sync.Mutex
}

func (rcv *neighbours) Connect(n string) error {
	neighbour, ok := rcv.nodes[n]
	if ok {
		return neighbour.Connect()
	}
	return fmt.Errorf("no such node: '%s", neighbour.Instance)
}
