package message

import (
	"bufio"
	"encoding/gob"
	"net"
	"strings"
	"time"
)

const (
	CreateTy   = "create"
	ReadTy     = "read"
	UpdateTy   = "update"
	DeleteTy   = "delete"
	pingTy     = "ping"
	pongTy     = "pong"
	msgTyDelim = '\n'
)

type Message interface {
	Type() string
}

func NewMessage(ty string) Message {
	switch ty {
	case CreateTy:
		return &Create{}
	case ReadTy:
		return &Read{}
	case UpdateTy:
		return &Update{}
	case DeleteTy:
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
	Id    uint64
	Bytes []byte
}

func (rcv *Create) Type() string {
	return CreateTy
}

// Read Message
type Read struct {
	Id uint64
}

func (rcv *Read) Type() string {
	return ReadTy
}

// Update message
type Update struct {
	Filename string
	Id       uint64
	Bytes    []byte
}

func (rcv *Update) Type() string {
	return UpdateTy
}

// Delete message
type Delete struct {
	Filename string
}

func (rcv *Delete) Type() string {
	return DeleteTy
}

func Send(msg Message, conn net.Conn) error {
	w := bufio.NewWriter(conn)
	enc := gob.NewEncoder(w)

	if _, err := w.WriteString(msg.Type() + string(msgTyDelim)); err != nil {
		return err
	}

	if err := enc.Encode(msg); err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}

	return nil
}

func Recv(conn net.Conn) (Message, error) {
	if err := conn.SetReadDeadline(time.Now().Add(1 * time.Nanosecond)); err != nil {
		return nil, err
	}

	r := bufio.NewReader(conn)
	msgType, err := r.ReadString(msgTyDelim)
	msgType = strings.Trim(msgType, string(msgTyDelim))

	if err != nil {
		return nil, err
	}

	msg := NewMessage(msgType)

	if err := gob.NewDecoder(r).Decode(msg); err != nil {
		return nil, err
	}

	return msg, nil
}
