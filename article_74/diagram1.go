package main

import (
	"log"

	"github.com/blushft/go-diagrams/diagram"
)

func main() {
	// inicializace objektu představujícího diagram
	diagram1, err := diagram.New(diagram.Label("Diagram #1"), diagram.Filename("diagram1"))

	// kontrola konstrukce objektu
	if err != nil {
		log.Fatal(err)
	}

	// vykreslení diagramu
	err = diagram1.Render()

	// kontrola, zda bylo vykreslení provedeno bez chyby
	if err != nil {
		log.Fatal(err)
	}
}
