// Seriál "Programovací jazyk Go"
//
// Třicátá první část
//    Sledování vybraných metrik služeb naprogramovaných v jazyku Go
//    https://www.root.cz/clanky/sledovani-vybranych-metrik-sluzeb-naprogramovanych-v-jazyku-go/
//
// Demonstrační příklad číslo 11:
//    Čítač atomicky zvyšovaný ve více vláknech.

package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var cnt int32 = 0

	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt32(&cnt, 1)
		}()
	}
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("%d\n", cnt)
}
