// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z devadesáté části:
//    https://github.com/tisnik/go-root/blob/master/article_90/README.md

package main

import (
	"fmt"
	"github.com/trustmaster/goflow"
	"strings"
)

// První typ uzlu v síti
type Printer struct {
	Message <-chan string
}

// Zpracování vstupních dat posílaných přes kanál
func (c *Printer) Process() {
	for message := range c.Message {
		fmt.Println(message)
	}
}

// Druhý typ uzlu v síti
type Converter struct {
	// vstupní port
	In <-chan string

	// první výstupní port
	Out1 chan<- string

	// druhý výstupní port
	Out2 chan<- string
}

// Zpracování vstupních dat posílaných přes kanál
// a poslání výsledku na výstupní kanály
func (c *Converter) Process() {
	for message := range c.In {
		converted := "to upper case: " + strings.ToUpper(message)
		c.Out1 <- converted

		converted = "to lower case: " + strings.ToLower(message)
		c.Out2 <- converted
	}
}

// Definice grafu sítě s aplikací
func NewFlowApp() *goflow.Graph {
	// konstrukce grafu
	n := goflow.NewGraph()

	// přidání uzlu s procesem do grafu
	n.Add("converter", new(Converter))

	// přidání uzlu s procesem do grafu
	n.Add("printer1", new(Printer))

	// přidání uzlu s procesem do grafu
	n.Add("printer2", new(Printer))

	// propojení dvou uzlů
	n.Connect("converter", "Out1", "printer1", "Message")

	// propojení dvou uzlů
	n.Connect("converter", "Out2", "printer2", "Message")

	// propojení uzlu s procesem
	n.MapInPort("Input", "converter", "In")

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
