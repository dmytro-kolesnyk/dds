package main

import (
	"log"

	"github.com/dmytro-kolesnyk/dds/discovery"
	"github.com/dmytro-kolesnyk/dds/node"
	"github.com/google/uuid"
)

func main() {
	discoverer := discovery.NewDiscovery(
		uuid.New().String(),
		"_dds._tcp",
		"local.",
		3451,
	)

	neighbours := make(chan node.Node)
	if err := discoverer.Start(neighbours); err != nil {
		log.Fatalln(err)
	}
	defer discoverer.Stop()

	for n := range neighbours {
		if n.Instance != discoverer.Instance {
			log.Println(n)
		}
	}

	select {}
}
