// Seriál "Programovací jazyk Go"
//
// Třicátá šestá část
//    Kooperace mezi kódem psaným v Go a C: cgo
//    https://www.root.cz/clanky/kooperace-mezi-kodem-psanym-v-go-a-c-cgo/

package main

// #include <math.h>
import "C"
import "fmt"

func main() {
	x := 3.1415927 / 6.0
	y := float32(C.sinf(C.float(x)))
	fmt.Printf("%v\n", y)
}
