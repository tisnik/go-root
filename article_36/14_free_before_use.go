// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá šestá část
//    Kooperace mezi kódem psaným v Go a C: cgo
//    https://www.root.cz/clanky/kooperace-mezi-kodem-psanym-v-go-a-c-cgo/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Demonstrační příklad číslo 14:
//    Uvolnění paměti před jejím použitím.
//
// Seznam demonstračních příkladů ze třicáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_36/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_36/14_free_before_use.html
//

package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

func main() {
	cs := C.CString("Hello world!")
	C.free(unsafe.Pointer(cs))
	C.puts(cs)
}
