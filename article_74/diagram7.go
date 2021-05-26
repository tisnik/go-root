package main

import (
	"fmt"
	"log"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/apps"
	"github.com/blushft/go-diagrams/nodes/gcp"
	"github.com/blushft/go-diagrams/nodes/generic"
)

const ProxiesCount = 3

func main() {
	// inicializace objektu představujícího diagram
	diagram7, err := diagram.New(diagram.Label("Diagram #7"), diagram.Filename("diagram7"))

	// kontrola konstrukce objektu
	if err != nil {
		log.Fatal(err)
	}

	// deklarace uzlů v diagramu
	inet := apps.Network.Internet().Label("Internet")

	proxies := make([]*diagram.Node, ProxiesCount)

	for i := 0; i < ProxiesCount; i++ {
		label := fmt.Sprintf("Proxy #%d", i+1)
		proxies[i] = generic.Network.Firewall().Label(label)
	}
	router := gcp.Network.Router().Label("Router")

	diagram7.Add(inet)
	diagram7.Add(router)

	diagram7.Group(diagram.NewGroup("Proxies").
		Add(proxies...).
		ConnectAllFrom(inet.ID()).
		ConnectAllTo(router.ID()),
	)

	// vykreslení diagramu
	err = diagram7.Render()

	// kontrola, zda bylo vykreslení provedeno bez chyby
	if err != nil {
		log.Fatal(err)
	}
}
