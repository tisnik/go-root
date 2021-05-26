package main

import (
	"log"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/apps"
)

func main() {
	// inicializace objektu představujícího diagram
	diagram2, err := diagram.New(diagram.Label("Diagram #2"), diagram.Filename("diagram2"))

	// kontrola konstrukce objektu
	if err != nil {
		log.Fatal(err)
	}

	// deklarace uzlů v diagramu
	inet := apps.Network.Internet().Label("Internet")
	proxy := apps.Network.Caddy().Label("3scale")

	// propojení uzlů v diagramu
	diagram2.Connect(inet, proxy)

	// vykreslení diagramu
	err = diagram2.Render()

	// kontrola, zda bylo vykreslení provedeno bez chyby
	if err != nil {
		log.Fatal(err)
	}
}
