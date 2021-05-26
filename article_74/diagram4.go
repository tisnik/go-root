package main

import (
	"log"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/apps"
	"github.com/blushft/go-diagrams/nodes/gcp"
)

func main() {
	// inicializace objektu představujícího diagram
	diagram4, err := diagram.New(diagram.Label("Diagram #4"), diagram.Filename("diagram4"))

	// kontrola konstrukce objektu
	if err != nil {
		log.Fatal(err)
	}

	// deklarace uzlů v diagramu
	inet := apps.Network.Internet().Label("Internet")
	proxy := apps.Network.Caddy().Label("3scale")
	router := gcp.Network.Router().Label("Router")

	diagram4.Group(diagram.NewGroup("Wild west").Add(inet).Add(proxy))

	// propojení uzlů v diagramu
	diagram4.Connect(inet, proxy)
	diagram4.Connect(proxy, router)

	// vykreslení diagramu
	err = diagram4.Render()

	// kontrola, zda bylo vykreslení provedeno bez chyby
	if err != nil {
		log.Fatal(err)
	}
}
