package connection

import (
	"bufio"
	"encoding/gob"
	"net"
	"strings"

	"github.com/dmytro-kolesnyk/dds/message"
)

type Connection struct {
	rw *bufio.ReadWriter
}

func NewConnection(conn net.Conn) *Connection {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	return &Connection{rw}
}

func (rcv *Connection) Send(msg message.Message) error {
	enc := gob.NewEncoder(rcv.rw.Writer)

	if _, err := rcv.rw.Writer.WriteString(msg.Type() + string(message.Delim)); err != nil {
		return err
	}

	if err := enc.Encode(msg); err != nil {
		return err
	}

	if err := rcv.rw.Writer.Flush(); err != nil {
		return err
	}

	return nil
}

func (rcv *Connection) RecvMsgType() (string, error) {
	cmd, err := rcv.rw.Reader.ReadString(message.Delim)
	cmd = strings.Trim(cmd, string(message.Delim))
	return cmd, err
}

func (rcv *Connection) Recv(msg message.Message) error {
	return gob.NewDecoder(rcv.rw.Reader).Decode(msg)
}
