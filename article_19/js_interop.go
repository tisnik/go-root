package main

import (
	"syscall/js"
)

func PrintHello(inputs []js.Value) {
	window := js.Global()
	document := window.Get("document")
	element := document.Call("getElementById", "header")

	element.Set("innerHTML", "Hello from Go")
}

func main() {
	println("started")
	c := make(chan bool)
	js.Global().Set("printHello", js.NewCallback(PrintHello))
	<-c
	println("finished")
}
