package app

import (
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/dmytro-kolesnyk/dds/connection"
	"github.com/dmytro-kolesnyk/dds/listener"
	"github.com/dmytro-kolesnyk/dds/message"
)

const (
	Port = ":3541"
)

type Server struct {
	*listener.Listener
	conns []connection.Connection
}

func NewServer() *Server {
	l := listener.NewListener()

	l.AddHandler(&message.Create{}, handleCreate)
	l.AddHandler(&message.Read{}, handleRead)
	l.AddHandler(&message.Update{}, handleUpdate)
	l.AddHandler(&message.Delete{}, handleDelete)

	return &Server{l, nil}
}

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

type Client struct{}

func (rcv *Client) Open(addr string, port string) (*connection.Connection, error) {
	log.Println("connect to", addr, "...")
	conn, err := net.Dial("tcp", addr+port)

	if err != nil {
		return nil, err
	}

	return connection.NewConnection(conn), nil
}

func (rcv *Client) Talk(addr string, port string) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	read := func() message.Message {
		return &message.Read{
			Id: 0xfeed_feed_0000_0000 | r.Uint64()%0xffff_ffff,
		}
	}()

	create := func() message.Message {
		bytes := make([]byte, r.Int()%0xff)
		//bytes := make([]byte, 4096)
		for i := 0; i < len(bytes); i++ {
			bytes[i] = byte(r.Intn(0xff))
		}
		return &message.Create{
			Id:    0xdefe_ca7e_0000_0000 | r.Uint64()%0xffff_ffff,
			Bytes: bytes,
		}
	}()

	update := func() message.Message {
		bytes := make([]byte, r.Int()%0xff)
		//bytes := make([]byte, 4096)
		for i := 0; i < len(bytes); i++ {
			bytes[i] = byte(r.Intn(0xff))
		}
		return &message.Update{
			Filename: "updateFilename.file",
			Id:       0xdefa_ce17_0000_0000 | r.Uint64()%0xffff_ffff,
			Bytes:    bytes,
		}
	}()

	del := func() message.Message {
		return &message.Delete{
			Filename: "deleteFilename.file",
		}
	}()

	msgs := []message.Message{create, read, update, del}

	s, err := rcv.Open(addr, port)
	if err != nil {
		return err
	}

	for {
		msg := msgs[r.Int()%len(msgs)]
		log.Printf("send %s message", msg.Type())
		if err := s.Send(msg); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * time.Duration(100+r.Int()%2400))
	}
}

type App interface {
	Start() error
	Stop() error
}

type Daemon struct {
	S *Server
	C *Client
}

func NewDaemon() App {
	return &Daemon{}
}

func (rcv *Daemon) startServer(port string) error {
	rcv.S = NewServer()
	if err := rcv.S.Listen(port); err != nil {
		return err
	}
	return nil
}

func (rcv *Daemon) startClient(addr string, port string) error {
	rcv.C = &Client{}
	if err := rcv.C.Talk(addr, port); err != nil {
		return err
	}
	return nil
}

func (rcv *Daemon) Start() error {
	connect := os.Getenv("CONNECT")

	if len(connect) > 0 {
		return rcv.startClient(connect, Port)
	}

	return rcv.startServer(Port)
}

func (rcv *Daemon) Stop() error {
	log.Println("Stopped")
	return nil
}
