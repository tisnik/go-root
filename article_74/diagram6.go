package main

import (
	"log"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/apps"
	"github.com/blushft/go-diagrams/nodes/gcp"
	"github.com/blushft/go-diagrams/nodes/generic"
)

func main() {
	// inicializace objektu představujícího diagram
	diagram6, err := diagram.New(diagram.Label("Diagram #6"), diagram.Filename("diagram6"))

	// kontrola konstrukce objektu
	if err != nil {
		log.Fatal(err)
	}

	// deklarace uzlů v diagramu
	inet := apps.Network.Internet().Label("Internet")
	proxy1 := generic.Network.Firewall().Label("3scale")
	proxy2 := generic.Network.Firewall().Label("3scale")
	proxy3 := generic.Network.Firewall().Label("3scale")
	router := gcp.Network.Router().Label("Router")

	diagram6.Connect(inet, proxy1)
	diagram6.Connect(inet, proxy2)
	diagram6.Connect(inet, proxy3)

	diagram6.Connect(proxy1, router)
	diagram6.Connect(proxy2, router)
	diagram6.Connect(proxy3, router)

	// vykreslení diagramu
	err = diagram6.Render()

	// kontrola, zda bylo vykreslení provedeno bez chyby
	if err != nil {
		log.Fatal(err)
	}
}
