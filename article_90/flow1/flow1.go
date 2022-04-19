package main

import (
	"fmt"
	"github.com/trustmaster/goflow"
)

// Jediný typ uzlu v síti
type Printer struct {
	// vstupní port
	Message <-chan string
}

// Zpracování vstupních dat posílaných přes kanál
func (c *Printer) Process() {
	for message := range c.Message {
		fmt.Println(message)
	}
}

// Definice grafu sítě s aplikací
func NewFlowApp() *goflow.Graph {
	// konstrukce grafu
	n := goflow.NewGraph()

	// přidání uzlu s procesem do grafu
	n.Add("printer", new(Printer))

	// propojení uzlu s procesem
	n.MapInPort("Input", "printer", "Message")

	// výsledný graf
	return n
}

func main() {
	// vytvoření sítě
	net := NewFlowApp()

	// konstrukce kanálu
	channel := make(chan string)
	net.SetInPort("Input", channel)

	// spuštění sítě
	// (kanál bude použit pro čekání na její ukončení)
	wait := goflow.Run(net)

	// poslání několika zpráv do sítě
	channel <- "Foo"
	channel <- "Bar"
	channel <- "Baz"

	// žádost o ukončení činnosti sítě
	close(channel)

	// čekání na ukončení činnosti sítě
	<-wait
}
