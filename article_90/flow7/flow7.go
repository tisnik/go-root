package main

import (
	"fmt"
	"github.com/trustmaster/goflow"
	"strings"
	"time"
	"unicode"
)

// Výchozí kapacita bufferů
const BufferCapacity = 100

// První typ uzlu v síti
type Printer struct {
	Message <-chan string
}

// Zpracování vstupních dat posílaných přes kanál
func (c *Printer) Process() {
	for message := range c.Message {
		fmt.Println(message)
		if unicode.IsLower([]rune(message)[0]) {
			time.Sleep(1 * time.Second)
		}
	}
}

// Druhý typ uzlu v síti
type Splitter struct {
	// vstupní port
	In <-chan string

	// první výstupní port
	Out1 chan<- string

	// druhý výstupní port
	Out2 chan<- string
}

// Zpracování vstupních dat posílaných přes kanál
func (c *Splitter) Process() {
	for message := range c.In {
		c.Out1 <- message
		c.Out2 <- message
	}
}

// Třetí typ uzlu v síti
type Converter1 struct {
	// vstupní port
	In <-chan string

	// výstupní port
	Out chan<- string
}

// Zpracování vstupních dat posílaných přes kanál
// a poslání výsledku na výstupní kanály
func (c *Converter1) Process() {
	for message := range c.In {
		converted := strings.ToUpper(message)
		c.Out <- converted
	}
}

// Čtvrtý typ uzlu v síti
type Converter2 struct {
	// vstupní port
	In <-chan string

	// výstupní port
	Out chan<- string
}

// Zpracování vstupních dat posílaných přes kanál
// a poslání výsledku na výstupní kanály
func (c *Converter2) Process() {
	for message := range c.In {
		converted := strings.ToLower(message)
		c.Out <- converted
	}
}

// Definice grafu sítě s aplikací
func NewFlowApp() *goflow.Graph {
	// konstrukce grafu
	n := goflow.NewGraph()

	// přidání uzlu s procesem do grafu
	n.Add("splitter", new(Splitter))

	// přidání uzlu s procesem do grafu
	n.Add("converter1", new(Converter1))

	// přidání uzlu s procesem do grafu
	n.Add("converter2", new(Converter2))

	// přidání uzlu s procesem do grafu
	n.Add("printer1", new(Printer))

	// přidání uzlu s procesem do grafu
	n.Add("printer2", new(Printer))

	// propojení dvou uzlů
	n.ConnectBuf("splitter", "Out1", "converter1", "In", BufferCapacity)

	// propojení dvou uzlů
	n.ConnectBuf("splitter", "Out2", "converter2", "In", BufferCapacity)

	// propojení dvou uzlů
	n.ConnectBuf("converter1", "Out", "printer1", "Message", BufferCapacity)

	// propojení dvou uzlů
	n.ConnectBuf("converter2", "Out", "printer2", "Message", BufferCapacity)

	// propojení uzlu s procesem
	n.MapInPort("Input", "splitter", "In")

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
	for i := 0; i < 10; i++ {
		channel <- "Foo"
		channel <- "Bar"
		channel <- "Baz"
	}

	// žádost o ukončení činnosti sítě
	close(channel)

	// čekání na ukončení činnosti sítě
	<-wait
}
