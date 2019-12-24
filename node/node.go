package node

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/dmytro-kolesnyk/dds/message"
)

type Node struct {
	Instance string
	Addr     net.IP
	Port     int
	Conn     net.Conn
}

func (rcv *Node) Connect() (err error) {
	log.Printf("connecting to %s:%d\n", rcv.Addr, rcv.Port)
	if rcv.Conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", rcv.Addr, rcv.Port)); err != nil {
		return err
	}

	if err = rcv.Conn.(*net.TCPConn).SetKeepAlive(true); err != nil {
		return err
	}

	return
}

func (rcv *Node) Start() error {
	return nil
}

/*
func PingPong() error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ping := &message.Ping{
		Xid: r.Uint32() % 0x0fff_ffff,
	}

	go func() {

	}()
	return nil
}
*/

func (rcv *Node) Send(msg message.Message) error {
	return message.Send(msg, rcv.Conn)
}

func (rcv *Node) Recv() (message.Message, error) {
	return message.Recv(rcv.Conn)
}

//
// [TODO] Mock function
//
func (rcv *Node) Talk() error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ping := func() message.Message {
		return &message.Ping{
			Xid: r.Uint32() % 0x0fff_ffff,
		}
	}
	/*
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

		messages := []message.Message{create, read, update, del}

		for {
			msg := messages[r.Int()%len(messages)]
			//log.Printf("send %s message to %s:%d\n", msg.Type(), rcv.Addr, rcv.Port)
			if err := rcv.Send(msg); err != nil {
				return err
			}
			time.Sleep(time.Millisecond * time.Duration(100+r.Int()%2400))
		}

	*/
	//go func() {
	for {
		request := ping()

		if err := rcv.Send(request); err != nil {
			log.Printf("error during sending '%s' request: %s\n", request.Type(), err)
			return err
		}

		log.Printf(
			"ping (xid: %d) sent to %s\n",
			request.(*message.Ping).Xid,
			rcv.Conn.RemoteAddr(),
		)

		if response, err := rcv.Recv(); err != nil {
			log.Printf("error during sending '%s' request: %s\n", request.Type(), err)
			return err
		} else {
			if response.(*message.Pong).Xid != request.(*message.Ping).Xid+1 {
				return fmt.Errorf("invalid Pong answer")
			}
		}

		//time.Sleep(time.Millisecond * time.Duration(100+r.Int()%2400))
		time.Sleep(time.Second * 5)
	}
}
