// Seriál "Programovací jazyk Go"
//
// Dvacátá devátá část
//    Trasování a profilování aplikací naprogramovaných v Go
//    https://www.root.cz/clanky/trasovani-a-profilovani-aplikaci-naprogramovanych-v-go/

package main

// #include <stdlib.h>
import "C"

func main() {
	C.srand(C.uint(42))

	for i := 1; i <= 10; i++ {
		x := C.random()
		println(x)
	}
}
