package discovery

import (
	"context"
	"fmt"

	"github.com/dmytro-kolesnyk/dds/common/logger"
	"github.com/dmytro-kolesnyk/dds/node"
	"github.com/grandcat/zeroconf"
)

type Discovery struct {
	Instance       string
	service        string
	domain         string
	port           int
	neighbours     chan *zeroconf.ServiceEntry
	logger         *logger.Logger
	server         *zeroconf.Server
	resolver       *zeroconf.Resolver
	resolverCtx    context.Context
	resolverCancel context.CancelFunc
}

func NewDiscovery(instance, service, domain string, port int) *Discovery {
	return &Discovery{
		Instance:   instance,
		service:    service,
		domain:     domain,
		port:       port,
		neighbours: make(chan *zeroconf.ServiceEntry),
		logger:     logger.NewLogger(&Discovery{}),
	}
}

func (rcv *Discovery) startServer() (err error) {
	rcv.server, err = zeroconf.Register(
		rcv.Instance,
		rcv.service,
		rcv.domain,
		rcv.port,
		nil,
		nil, // [TODO] get from config.yaml
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
	rcv.logger.Info("looking for neighbours")
	if err := rcv.startServer(); err != nil {
		return err
	}

	if err := rcv.startResolver(); err != nil {
		return err
	}

	//neighbours := make(chan *zeroconf.ServiceEntry)
	if err := rcv.browse(rcv.neighbours); err != nil {
		return err
	}

	go func() {
		for n := range rcv.neighbours {
			if n.Instance != rcv.Instance {
				neighbour := node.NewNode(
					n.Instance,
					n.AddrIPv4[0], // [FIXME] should select neighbour IP/IPv6 addr in more intelligent way
					n.Port,
				)
				nodes <- neighbour
				rcv.logger.Info(fmt.Sprintf("new neighbour: %s", neighbour))
			}
		}
	}()

	return nil
}

func (rcv *Discovery) Stop() {
	rcv.logger.Info("stopping mDNS server")
	rcv.server.Shutdown()
	rcv.logger.Info("stopping mDNS resolver")
	<-rcv.resolverCtx.Done()
	rcv.resolverCancel()
	close(rcv.neighbours)
}
