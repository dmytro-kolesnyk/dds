package communicationServer

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/dmytro-kolesnyk/dds/common/logger"

	"github.com/dmytro-kolesnyk/dds/common/conf/models"

	"github.com/dmytro-kolesnyk/dds/listener"

	"github.com/dmytro-kolesnyk/dds/discovery"
	"github.com/dmytro-kolesnyk/dds/node"
	"github.com/google/uuid"

	"github.com/dmytro-kolesnyk/dds/message"
)

func handleCreate(req message.Message, log *logger.Logger) message.Message {
	//data := msg.(*message.Create)
	//log.Printf("%#v\n", data)
	return nil
}

func handleRead(req message.Message, log *logger.Logger) message.Message {
	//data := msg.(*message.Read)
	//log.Printf("%#v\n", data)
	return nil
}

func handleUpdate(req message.Message, log *logger.Logger) message.Message {
	//data := msg.(*message.Update)
	//log.Printf("%#v\n", data)
	return nil
}

func handleDelete(req message.Message, log *logger.Logger) message.Message {
	//data := msg.(*message.Delete)
	//log.Printf("%#v\n", data)
	return nil
}

func handlePing(req message.Message, log *logger.Logger) message.Message {
	ping := req.(*message.Ping)
	pong := &message.Pong{Xid: ping.Xid + 1}

	log.Debug(fmt.Sprintf("%s(xid=%d) for %s(xid=%d)", ping.Type(), ping.Xid, pong.Type(), pong.Xid))

	return pong
}

type CommunicationServer struct {
	listener   *listener.Listener
	discoverer *discovery.Discovery
	nodes      sync.Map
	port       int
	logger     *logger.Logger
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
		logger.NewLogger(&CommunicationServer{}),
	}
}

func (rcv *CommunicationServer) Start() error {
	neighbours := make(chan *node.Node)
	dead := make(chan string)

	if err := rcv.discoverer.Start(neighbours); err != nil {
		rcv.logger.Error(err)
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
					if value.(*node.Node).Conn != nil && value.(*node.Node).Queue != nil {
						return
					}
					rcv.logger.Info("talking with", value)
					if err := value.(*node.Node).Start(); err != nil {
						rcv.logger.Error(key, err)
						dead <- key.(string)
						close(value.(*node.Node).Queue)
						if err := value.(*node.Node).Conn.Close(); err != nil {
							rcv.logger.Error(key, err)
						}
						return
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

	// [TODO] get IP from config.yaml
	if err := rcv.listener.Start(fmt.Sprintf(":%d", rcv.port)); err != nil {
		return err
	}

	return nil
}

func (rcv *CommunicationServer) Stop() error {
	rcv.logger.Info("stopping")
	rcv.logger.Info("cleaning up neighbours")
	rcv.nodes.Range(func(key, value interface{}) bool {
		if err := value.(*node.Node).Conn.Close(); err != nil {
			rcv.logger.Error(err)
		}
		rcv.nodes.Delete(key)

		return true
	})
	rcv.discoverer.Stop()

	if err := rcv.listener.Stop(); err != nil {
		return err
	}
	return nil
}
