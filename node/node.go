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
	log.Println("connecting to", rcv.Addr)
	rcv.Conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", rcv.Addr, rcv.Port))

	return
}

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
		log.Printf("send %s message to %s:%d\n", msg.Type(), rcv.Addr, rcv.Port)
		if err := rcv.Send(msg); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * time.Duration(100+r.Int()%2400))
	}
}
