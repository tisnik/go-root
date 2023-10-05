// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam příkladů ze sedmdesáté čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_74/README.md

package main

import (
	"log"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/apps"
	"github.com/blushft/go-diagrams/nodes/gcp"
)

func main() {
	// inicializace objektu představujícího diagram
	diagram5, err := diagram.New(diagram.Label("Diagram #5"), diagram.Filename("diagram5"))

	// kontrola konstrukce objektu
	if err != nil {
		log.Fatal(err)
	}

	// deklarace uzlů v diagramu
	inet := apps.Network.Internet().Label("Internet")
	proxy := apps.Network.Caddy().Label("3scale")
	router := gcp.Network.Router().Label("Router")

	diagram5.Group(diagram.NewGroup("Wild west").Label("Wild west").Add(inet).Add(proxy))

	// propojení uzlů v diagramu
	diagram5.Connect(inet, proxy)
	diagram5.Connect(proxy, router)

	// vykreslení diagramu
	err = diagram5.Render()

	// kontrola, zda bylo vykreslení provedeno bez chyby
	if err != nil {
		log.Fatal(err)
	}
}
