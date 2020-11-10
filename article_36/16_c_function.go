// Seriál "Programovací jazyk Go"
//
// Třicátá šestá část
//    Kooperace mezi kódem psaným v Go a C: cgo
//    https://www.root.cz/clanky/kooperace-mezi-kodem-psanym-v-go-a-c-cgo/

package main

// int add(int a, int b) {
//     return a+b;
// }
import "C"
import "fmt"

func main() {
	x := C.add(C.int(1), C.int(2))
	fmt.Printf("result=%d\n", x)
}
