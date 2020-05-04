// Seriál "Programovací jazyk Go"
//
// Devatenáctá část
//     Využití WebAssembly z programovacího jazyka Go
//     https://www.root.cz/clanky/vyuziti-webassembly-z-programovaciho-jazyka-go/
//
// Demonstrační příklad:
//     Kreslení na Canvas.

package main

import (
	"syscall/js"
)

func main() {
	const CanvasWidth = 256
	const CanvasHeight = 256

	println("started")

	window := js.Global()
	document := window.Get("document")

	canvas := document.Call("createElement", "canvas")
	canvas.Set("height", CanvasWidth)
	canvas.Set("width", CanvasHeight)
	document.Get("body").Call("appendChild", canvas)

	context2d := canvas.Call("getContext", "2d")
	context2d.Set("fillStyle", "#c0c0c0")
	context2d.Call("fillRect", 0, 0, CanvasWidth, CanvasHeight)

	context2d.Set("fillStyle", "yellow")
	context2d.Call("fillRect", 10, 10, CanvasWidth-20, CanvasHeight-20)

	println("finished")
}
