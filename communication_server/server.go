package communicationServer

import (
	"github.com/dmytro-kolesnyk/dds/listener"
	"github.com/dmytro-kolesnyk/dds/message"
	"github.com/dmytro-kolesnyk/dds/node"
)

type server struct {
	*listener.Listener
	Nodes []*node.Node
}

func newServer() *server {
	l := listener.NewListener()

	l.AddHandler(&message.Create{}, handleCreate)
	l.AddHandler(&message.Read{}, handleRead)
	l.AddHandler(&message.Update{}, handleUpdate)
	l.AddHandler(&message.Delete{}, handleDelete)

	return &server{l, nil}
}
