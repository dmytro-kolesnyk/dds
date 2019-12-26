package message

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"strings"
)

const (
	createTy   = "create"
	readTy     = "read"
	updateTy   = "update"
	deleteTy   = "delete"
	pingTy     = "ping"
	pongTy     = "pong"
	msgTyDelim = '\n'
)

type Message interface {
	Type() string
}

func NewMessage(ty string) Message {
	switch ty {
	case createTy:
		return &Create{}
	case readTy:
		return &Read{}
	case updateTy:
		return &Update{}
	case deleteTy:
		return &Delete{}
	case pingTy:
		return &Ping{}
	case pongTy:
		return &Pong{}
	}

	return nil
}

// Ping Message
type Ping struct {
	Xid uint32
}

func (rcv *Ping) Type() string {
	return pingTy
}

// Pong Message
type Pong struct {
	Xid uint32
}

func (rcv *Pong) Type() string {
	return pongTy
}

// Create Message
type Create struct {
	Xid   uint64
	Id    uint64
	Bytes []byte
}

func (rcv *Create) Type() string {
	return createTy
}

// Read Message
type Read struct {
	Xid uint64
	Id  uint64
}

func (rcv *Read) Type() string {
	return readTy
}

// Update message
type Update struct {
	Xid      uint64
	Filename string
	Id       uint64
	Bytes    []byte
}

func (rcv *Update) Type() string {
	return updateTy
}

// Delete message
type Delete struct {
	Xid      uint64
	Filename string
}

func (rcv *Delete) Type() string {
	return deleteTy
}

func Send(msg Message, conn net.Conn) error {
	w := bufio.NewWriter(conn)
	enc := gob.NewEncoder(w)

	if _, err := w.WriteString(msg.Type() + string(msgTyDelim)); err != nil {
		return fmt.Errorf("send '%s' to %s: %s", msg.Type(), conn.RemoteAddr(), err)
	}

	if err := enc.Encode(msg); err != nil {
		return fmt.Errorf("send '%s' to %s: %s", msg.Type(), conn.RemoteAddr(), err)
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf("send '%s' to %s: %s", msg.Type(), conn.RemoteAddr(), err)
	}

	return nil
}

func Recv(conn net.Conn) (Message, error) {
	// [TODO] move timeouts
	//if err := conn.SetReadDeadline(time.Now().Add(1 * time.Nanosecond)); err != nil {
	//	return nil, err
	//}

	r := bufio.NewReader(conn)
	msgType, err := r.ReadString(msgTyDelim)
	msgType = strings.Trim(msgType, string(msgTyDelim))

	if err != nil {
		return nil, fmt.Errorf("recv '%s' from %s: %s", msgType, conn.RemoteAddr(), err)
	}

	msg := NewMessage(msgType)
	if err := gob.NewDecoder(r).Decode(msg); err != nil {
		return nil, fmt.Errorf("recv '%s' from %s: %s", msgType, conn.RemoteAddr(), err)
	}

	return msg, nil
}
