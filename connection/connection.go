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

// [TODO] obsolete
func (rcv *Connection) RecvMsgType() (string, error) {
	cmd, err := rcv.rw.Reader.ReadString(message.Delim)
	cmd = strings.Trim(cmd, string(message.Delim))
	return cmd, err
}

// [TODO] obsolete
func (rcv *Connection) Recv(msg message.Message) error {
	return gob.NewDecoder(rcv.rw.Reader).Decode(msg)
}

func (rcv *Connection) RecvFull() (message.Message, error) {
	msgType, err := rcv.rw.Reader.ReadString(message.Delim)
	msgType = strings.Trim(msgType, string(message.Delim))

	if err != nil {
		return nil, err
	}

	msg := message.NewMessage(msgType)

	if err := gob.NewDecoder(rcv.rw.Reader).Decode(msg); err != nil {
		return nil, err
	}

	return msg, nil
}
