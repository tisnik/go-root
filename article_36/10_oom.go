// Seriál "Programovací jazyk Go"
//
// Třicátá šestá část
//    Kooperace mezi kódem psaným v Go a C: cgo
//    https://www.root.cz/clanky/kooperace-mezi-kodem-psanym-v-go-a-c-cgo/

package main

// #include <stdio.h>
import "C"

func main() {
	for {
		cs := C.CString("Hello world!")
		C.puts(cs)
	}
}
