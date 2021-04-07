// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá první část
//    Sledování vybraných metrik služeb naprogramovaných v jazyku Go
//    https://www.root.cz/clanky/sledovani-vybranych-metrik-sluzeb-naprogramovanych-v-jazyku-go/
//
// Demonstrační příklad číslo 9:
//    Čítač zvyšovaný ve více vláknech bez synchronizace.

package main

import (
	"fmt"
	"time"
)

func main() {
	var cnt int = 0

	for i := 0; i < 1000; i++ {
		go func() {
			cnt++
		}()
	}
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("%d\n", cnt)
}
