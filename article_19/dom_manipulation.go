// Seriál "Programovací jazyk Go"
//
// Devatenáctá část
//     Využití WebAssembly z programovacího jazyka Go
//     https://www.root.cz/clanky/vyuziti-webassembly-z-programovaciho-jazyka-go/
//
// Demonstrační příklad:
//     Manipulace s DOMem z jazyka Go.

package main

import (
	"syscall/js"
)

func main() {
	println("started")

	window := js.Global()
	document := window.Get("document")
	element := document.Call("getElementById", "header")

	element.Set("innerHTML", "foobar")

	println("finished")
}
