package communicationServer

import (
	"fmt"
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
	//data := msg.(*message.Create)
	//log.Printf("%#v\n", data)
}

func handleRead(msg message.Message) {
	//data := msg.(*message.Read)
	//log.Printf("%#v\n", data)
}

func handleUpdate(msg message.Message) {
	//data := msg.(*message.Update)
	//log.Printf("%#v\n", data)
}

func handleDelete(msg message.Message) {
	//data := msg.(*message.Delete)
	//log.Printf("%#v\n", data)
}

type CommunicationServer struct {
	listener   *listener.Listener
	discoverer *discovery.Discovery
	nodes      sync.Map
	port       int
	//nodes      map[string]*node.Node
}

func NewCommunicationServer(port int) *CommunicationServer {
	return &CommunicationServer{
		listener.NewListener(),
		discovery.NewDiscovery(
			uuid.New().String(),
			"_dds._tcp",
			"local.",
			port,
		),
		sync.Map{},
		port,
	}
}

func (rcv *CommunicationServer) Start() error {
	neighbours := make(chan *node.Node)
	dead := make(chan string)

	log.Println("looking for neighbours")
	if err := rcv.discoverer.Start(neighbours); err != nil {
		return err
	}

	go func() {
		for {
			select {
			case n := <-neighbours:
				rcv.nodes.Store(n.Instance, n)
			case instance := <-dead:
				rcv.nodes.Delete(instance)
				log.Printf("neighbour %s removed\n", instance)
			}
		}
	}()

	go func() {
		for {
			rcv.nodes.Range(func(key, value interface{}) bool {
				go func() {
					if value.(*node.Node).Conn != nil {
						return
					}
					log.Println("talking with", key)
					if err := value.(*node.Node).Connect(); err != nil {
						log.Println(key, err)
						dead <- key.(string)
						return
					}
					if err := value.(*node.Node).Talk(); err != nil {
						log.Println(key, err)
						dead <- key.(string)
					}
				}()
				return true
			})
			time.Sleep(1 * time.Second) // [BUG] sync issues
		}
	}()

	rcv.listener.AddHandler(&message.Create{}, handleCreate)
	rcv.listener.AddHandler(&message.Read{}, handleRead)
	rcv.listener.AddHandler(&message.Update{}, handleUpdate)
	rcv.listener.AddHandler(&message.Delete{}, handleDelete)

	if err := rcv.listener.Start(fmt.Sprintf(":%d", rcv.port)); err != nil {
		return err
	}

	return nil
}

func (rcv *CommunicationServer) Stop() error {
	rcv.discoverer.Stop()

	if err := rcv.listener.Stop(); err != nil {
		return err
	}
	return nil
}
