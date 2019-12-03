package app

type App interface {
	Start() error
	Stop() error
}

type Daemon struct {

}

func (rcv *Daemon) Start() error {
	panic("implement me")
}

func (rcv *Daemon) Stop() error {
	panic("implement me")
}



