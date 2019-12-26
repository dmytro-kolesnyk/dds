package listener

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/dmytro-kolesnyk/dds/common/logger"
	"github.com/dmytro-kolesnyk/dds/message"
)

type HandleFunc func(message.Message, *logger.Logger) message.Message

type Listener struct {
	listener net.Listener
	handler  map[string]HandleFunc
	mux      sync.RWMutex
	logger   *logger.Logger
}

func NewListener() *Listener {
	return &Listener{
		handler: map[string]HandleFunc{},
		logger:  logger.NewLogger(&Listener{}),
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

	rcv.logger.Info("listen on", rcv.listener.Addr())
	for {
		conn, err := rcv.listener.Accept()
		rcv.logger.Info("received connect from", conn.RemoteAddr())
		if err != nil {
			log.Println(err)
			continue
		}
		rcv.logger.Info("connected from", conn.RemoteAddr())
		go rcv.handle(conn)
	}
}

func (rcv *Listener) Stop() error {
	rcv.logger.Info("stopping")
	return rcv.listener.Close()
}

func (rcv *Listener) handle(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			rcv.logger.Error(err)
		}
	}()

	for {
		request, err := message.Recv(conn)

		if err != nil {
			rcv.logger.Error(err)
			return
		}

		rcv.mux.RLock()
		handler, ok := rcv.handler[request.Type()]
		rcv.mux.RUnlock()

		if !ok {
			rcv.logger.Warn(fmt.Sprintf("no handler for request '%s'", request.Type()))
			continue
		}

		rcv.logger.Info(fmt.Sprintf("request '%s' from %s received", request.Type(), conn.RemoteAddr()))
		if response := handler(request, rcv.logger); response != nil {
			if err := message.Send(response, conn); err != nil {
				rcv.logger.Warn(err)
			}
		}
	}
}
