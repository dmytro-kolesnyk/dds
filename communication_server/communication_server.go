package communicationServer

import (
	"fmt"
	"github.com/dmytro-kolesnyk/dds/common/conf/models"
	"log"
	"sync"
	"time"

	"github.com/dmytro-kolesnyk/dds/listener"

	"github.com/dmytro-kolesnyk/dds/discovery"
	"github.com/dmytro-kolesnyk/dds/node"
	"github.com/google/uuid"

	"github.com/dmytro-kolesnyk/dds/message"
)

func handleCreate(req message.Message) message.Message {
	//data := msg.(*message.Create)
	//log.Printf("%#v\n", data)
	return nil
}

func handleRead(req message.Message) message.Message {
	//data := msg.(*message.Read)
	//log.Printf("%#v\n", data)
	return nil
}

func handleUpdate(req message.Message) message.Message {
	//data := msg.(*message.Update)
	//log.Printf("%#v\n", data)
	return nil
}

func handleDelete(req message.Message) message.Message {
	//data := msg.(*message.Delete)
	//log.Printf("%#v\n", data)
	return nil
}

func handlePing(req message.Message) message.Message {
	ping := req.(*message.Ping)
	time.Sleep(5 * time.Nanosecond)
	log.Printf("pong (xid: %d) for ping (xid: %d)\n", ping.Xid+1, ping.Xid)
	return &message.Pong{
		Xid: ping.Xid + 1,
	}
}

type CommunicationServer struct {
	listener   *listener.Listener
	discoverer *discovery.Discovery
	nodes      sync.Map
	port       int
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
	rcv.listener.AddHandler(&message.Ping{}, handlePing)

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
