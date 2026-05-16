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
//    Manipulace s DOMem z jazyka Go.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_19/dom_manipulation.html

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
