package message

const (
	Delim    = '\n'
	CreateTy = "create"
	ReadTy   = "read"
	UpdateTy = "update"
	DeleteTy = "delete"
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
	}

	return nil
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
