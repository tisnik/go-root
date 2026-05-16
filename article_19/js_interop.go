// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devatenáctá část
//     Využití WebAssembly z programovacího jazyka Go
//     https://www.root.cz/clanky/vyuziti-webassembly-z-programovaciho-jazyka-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z devatenácté části:
//    https://github.com/tisnik/go-root/blob/master/article_19/README.md
//
// Demonstrační příklad:
//    Standardní balíček syscall/js.
//    Rozhraní mezi jazyky Go a JavaScript.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_19/js_interop.html

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
