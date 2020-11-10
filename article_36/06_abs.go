// Seriál "Programovací jazyk Go"
//
// Třicátá šestá část
//    Kooperace mezi kódem psaným v Go a C: cgo
//    https://www.root.cz/clanky/kooperace-mezi-kodem-psanym-v-go-a-c-cgo/

package main

// #include <stdlib.h>
import "C"
import "fmt"

func main() {
	x := -10
	y := C.abs(C.int(x))
	fmt.Printf("%v\n", y)
}
