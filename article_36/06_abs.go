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
// Seznam demonstračních příkladů ze třicáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_36/README.md

package main

// #include <stdlib.h>
import "C"
import "fmt"

func main() {
	x := -10
	y := C.abs(C.int(x))
	fmt.Printf("%v\n", y)
}
