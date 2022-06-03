// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devatenáctá část
//     Využití WebAssembly z programovacího jazyka Go
//     https://www.root.cz/clanky/vyuziti-webassembly-z-programovaciho-jazyka-go/
//
// Seznam příkladů z devatenácté části:
//    https://github.com/tisnik/go-root/blob/master/article_19/README.md
//
// Demonstrační příklad:
//     Manipulace s DOMem z jazyka Go.

package main

import (
	"fmt"
	"syscall/js"
)

func Factorial(n int64) int64 {
	switch {
	case n < 0:
		return 1
	case n == 0:
		return 1
	default:
		return n * Factorial(n-1)
	}
}

func main() {
	println("started")

	window := js.Global()
	document := window.Get("document")
	element := document.Call("getElementById", "header")

	element.Set("innerHTML", "foobar")

	for n := int64(0); n <= 10; n++ {
		f := Factorial(n)
		message := fmt.Sprintf("%2d! = %d", n, f)
		pre := document.Call("createElement", "pre")
		pre.Set("innerHTML", message)
		document.Get("body").Call("appendChild", pre)
	}

	println("finished")
}
