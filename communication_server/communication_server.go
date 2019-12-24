package communicationServer

import (
	"fmt"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/conf/models"
	"log"
	"sync"
	"time"

	"github.com/dmytro-kolesnyk/dds/listener"

	"github.com/dmytro-kolesnyk/dds/discovery"
	"github.com/dmytro-kolesnyk/dds/node"
	"github.com/google/uuid"

	"github.com/dmytro-kolesnyk/dds/message"
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
	listener   *listener.Listener
	discoverer *discovery.Discovery
	*neighbours
	port int
}

func NewCommunicationServer(config *models.Config) *CommunicationServer {
	port := config.CommunicationServer.Port
	return &CommunicationServer{
		listener.NewListener(),
		discovery.NewDiscovery(
			uuid.New().String(),
			"_dds._tcp",
			"local.",
			port,
		),
		&neighbours{make(map[string]*node.Node), sync.Mutex{}},
		port,
	}
}

func (rcv *CommunicationServer) Start() error {
	neighbours := make(chan *node.Node)
	if err := rcv.discoverer.Start(neighbours); err != nil {
		return err
	}

	log.Println("looking for neighbours")
	go func() {
		for n := range neighbours {
			rcv.mux.Lock()
			if _, ok := rcv.nodes[n.Instance]; !ok {
				rcv.nodes[n.Instance] = n
			}
			rcv.mux.Unlock()
			if err := rcv.neighbours.Connect(n.Instance); err != nil {
				log.Println(err)
			}
			log.Printf("new neighbour: %#+v\n", n)
		}
		log.Println("DISCOVERY STOPPED")
	}()

	// [TODO] Mock goroutine
	go func() {
		for {
			for _, n := range rcv.nodes {
				log.Println("neighbour:", n.Instance, n.Addr, n.Port)
				if err := n.Talk(); err != nil {
					log.Println(err)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()

	rcv.listener.AddHandler(&message.Create{}, handleCreate)
	rcv.listener.AddHandler(&message.Read{}, handleRead)
	rcv.listener.AddHandler(&message.Update{}, handleUpdate)
	rcv.listener.AddHandler(&message.Delete{}, handleDelete)

	if err := rcv.listener.Listen(fmt.Sprintf(":%d", rcv.port)); err != nil {
		return err
	}

	return nil
}

func (rcv *CommunicationServer) Stop() error {
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
