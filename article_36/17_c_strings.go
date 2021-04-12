// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá šestá část
//    Kooperace mezi kódem psaným v Go a C: cgo
//    https://www.root.cz/clanky/kooperace-mezi-kodem-psanym-v-go-a-c-cgo/

package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "fmt"

func main() {
	fmt.Print("? ")
	s := C.calloc(20, 1)
	charPtr := (*C.char)(s)
	C.fgets(charPtr, 20, C.stdin)
	fmt.Printf("result=%v\n", C.GoString(charPtr))
	C.free(s)
}
