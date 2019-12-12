package discovery

import (
	"context"

	"github.com/dmytro-kolesnyk/dds/node"
	"github.com/grandcat/zeroconf"
)

type Discovery struct {
	Instance       string
	service        string
	domain         string
	port           int
	server         *zeroconf.Server
	resolver       *zeroconf.Resolver
	resolverCtx    context.Context
	resolverCancel context.CancelFunc
}

func NewDiscovery(instance, service, domain string, port int) *Discovery {
	return &Discovery{
		Instance: instance,
		service:  service,
		domain:   domain,
		port:     port,
	}
}

func (rcv *Discovery) startServer() (err error) {
	rcv.server, err = zeroconf.Register(
		rcv.Instance,
		rcv.service,
		rcv.domain,
		rcv.port,
		nil,
		nil,
	)

	return
}

func (rcv *Discovery) startResolver() (err error) {
	rcv.resolver, err = zeroconf.NewResolver(nil)
	return err
}

func (rcv *Discovery) browse(neighbours chan *zeroconf.ServiceEntry) error {
	rcv.resolverCtx, rcv.resolverCancel = context.WithCancel(context.Background())
	if err := rcv.resolver.Browse(rcv.resolverCtx, rcv.service, rcv.domain, neighbours); err != nil {
		return err
	}

	return nil
}

func (rcv *Discovery) Start(nodes chan *node.Node) error {
	if err := rcv.startServer(); err != nil {
		return err
	}

	if err := rcv.startResolver(); err != nil {
		return err
	}

	neighbours := make(chan *zeroconf.ServiceEntry)
	if err := rcv.browse(neighbours); err != nil {
		return err
	}

	go func() {
		for n := range neighbours {
			if n.Instance != rcv.Instance {
				nodes <- &node.Node{
					Instance: n.Instance,
					Addr:     n.AddrIPv4[0], // [TODO] [BUG] [FIXME]
					Port:     n.Port,
				}
			}
		}
	}()

	return nil
}

func (rcv *Discovery) Stop() {
	rcv.server.Shutdown()
	<-rcv.resolverCtx.Done()
	rcv.resolverCancel()
}
