package listener

import (
	"encoding/gob"
	"log"
	"net"
	"sync"

	"github.com/dmytro-kolesnyk/dds/message"
)

type HandleFunc func(message.Message) message.Message

type Listener struct {
	listener net.Listener
	handler  map[string]HandleFunc
	mux      sync.RWMutex
}

func NewListener() *Listener {
	return &Listener{
		handler: map[string]HandleFunc{},
	}
}

func (rcv *Listener) AddHandler(m message.Message, f HandleFunc) {
	gob.Register(m)

	rcv.mux.Lock()
	rcv.handler[m.Type()] = f
	rcv.mux.Unlock()
}

func (rcv *Listener) Start(port string) error {
	var err error

	if rcv.listener, err = net.Listen("tcp", port); err != nil {
		return err
	}

	log.Println("listen on", rcv.listener.Addr())

	for {
		conn, err := rcv.listener.Accept()
		log.Println("received connect from", conn.RemoteAddr())
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("connected from", conn.RemoteAddr())
		go rcv.handle(conn)
	}
}

func (rcv *Listener) Stop() error {
	return rcv.listener.Close()
}

func (rcv *Listener) handle(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println(err)
		}
	}()

	for {
		request, err := message.Recv(conn)

		if err != nil {
			log.Printf("%s disconnected, %s\n", conn.RemoteAddr(), err)
			//if err == io.EOF {
			return
			//}
		}

		rcv.mux.RLock()
		handler, ok := rcv.handler[request.Type()]
		rcv.mux.RUnlock()

		if !ok {
			log.Printf("handler for message:'%s' is not registered\n", request.Type())
			continue
		}

		log.Printf("request '%s' from %s received\n", request.Type(), conn.RemoteAddr())
		response := handler(request)
		if response != nil {
			if err := message.Send(response, conn); err != nil {
				log.Println("error during response:", err)
			}
		}
	}
}
