// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá devátá část
//    Trasování a profilování aplikací naprogramovaných v Go
//    https://www.root.cz/clanky/trasovani-a-profilovani-aplikaci-naprogramovanych-v-go/
//
// Seznam příkladů z dvacáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_29/README.md

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
