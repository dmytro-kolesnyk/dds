package node

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/dmytro-kolesnyk/dds/common/logger"
	"github.com/dmytro-kolesnyk/dds/message"
)

type Node struct {
	Instance string
	Addr     net.IP
	Port     int
	Conn     net.Conn
	logger   *logger.Logger
	Queue    chan message.Message
}

func NewNode(instance string, addr net.IP, port int) *Node {
	return &Node{
		Instance: instance,
		Addr:     addr,
		Port:     port,
		logger:   logger.NewLogger(&Node{}),
	}
}

func (rcv *Node) String() string {
	return fmt.Sprintf("%s/%s", rcv.Instance, rcv.Endpoint())
}

func (rcv *Node) Endpoint() string {
	return fmt.Sprintf("%s:%d", rcv.Addr, rcv.Port)
}

func (rcv *Node) Connect() (err error) {
	rcv.logger.Info("connecting to", rcv.Endpoint())
	if rcv.Conn, err = net.Dial("tcp", rcv.Endpoint()); err != nil {
		return
	}

	return rcv.Conn.(*net.TCPConn).SetKeepAlive(true)
}

func (rcv *Node) Start() error {
	if err := rcv.Connect(); err != nil {
		return err
	}

	rcv.Queue = make(chan message.Message)

	go func() {
		for r := range rcv.Queue {
			if err := rcv.Send(r); err != nil {
				rcv.logger.Error(err)
				return
			}
		}
	}()

	//
	// [TODO] mock
	//
	go func() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for {
			request := &message.Ping{
				Xid: r.Uint32() % 0x0fff_ffff,
			}

			rcv.Queue <- request
			if response, err := rcv.Recv(); err != nil {
				rcv.logger.Error(err)
				return
			} else {
				if err := handlePong(request, response); err != nil {
					rcv.logger.Error(err)
				}
			}
			time.Sleep(time.Second * 5)
		}
	}()

	return nil
}

func handlePong(req message.Message, resp message.Message) error {
	if resp.(*message.Pong).Xid != req.(*message.Ping).Xid+1 {
		return fmt.Errorf("invalid '%s' response", resp.Type())
	}
	return nil
}

func (rcv *Node) Send(msg message.Message) error {
	rcv.logger.Info(fmt.Sprintf("sending '%s' request to %s", msg.Type(), rcv))
	return message.Send(msg, rcv.Conn)
}

func (rcv *Node) Recv() (message.Message, error) {
	resp, err := message.Recv(rcv.Conn)
	if err != nil {
		return nil, err
	}
	rcv.logger.Info(fmt.Sprintf("received '%s' response from %s", resp.Type(), rcv))

	return resp, nil
}

/*
func (rcv *Node) handle() {
	defer func() {
		if err := rcv.Conn.Close(); err != nil {
			rcv.logger.Error(err)
		}
	}()

}
*/
