// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá šestá část
//    Kooperace mezi kódem psaným v Go a C: cgo
//    https://www.root.cz/clanky/kooperace-mezi-kodem-psanym-v-go-a-c-cgo/

package main

// #include <math.h>
// #cgo LDFLAGS: -lm
import "C"
import "fmt"

func main() {
	x := 3.1415927 / 6.0
	y := float32(C.sinf(C.float(x)))
	fmt.Printf("%v\n", y)
}
