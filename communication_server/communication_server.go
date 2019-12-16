package communicationServer

import (
	"log"
	"time"

	"github.com/dmytro-kolesnyk/dds/message"
	"github.com/dmytro-kolesnyk/dds/node"
)

const (
	Port = ":3541" // [FIXME] take from config
)

func handleCreate(msg message.Message) {
	data := msg.(*message.Create)
	log.Printf("%#v\n", data)
}

func handleRead(msg message.Message) {
	data := msg.(*message.Read)
	log.Printf("%#v\n", data)
}

func handleUpdate(msg message.Message) {
	data := msg.(*message.Update)
	log.Printf("%#v\n", data)
}

func handleDelete(msg message.Message) {
	data := msg.(*message.Delete)
	log.Printf("%#v\n", data)
}

type CommunicationServer struct {
	*server
	*client
}

func NewCommunicationServer() *CommunicationServer {
	return &CommunicationServer{}
}

func (rcv *CommunicationServer) Start(port string, neighbours chan *node.Node) error {
	rcv.server = newServer()
	if err := rcv.Listen(port); err != nil {
		return err
	}

	go func() {
		for n := range neighbours {
			rcv.Nodes = append(rcv.Nodes, n)
			log.Printf("new neighbour: %#+v\n", n)
			time.Sleep(5 * time.Second)
		}
	}()

	//go func() {
	//
	//}()

	return nil
}

/*
func (rcv *CommunicationServer) startClient(addr string, port string) error {
	rcv.client = &client{}
	if err := rcv.Talk(addr, port); err != nil {
		return err
	}
	return nil
}
*/
