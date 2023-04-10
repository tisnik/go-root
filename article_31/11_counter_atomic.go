// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá první část
//    Sledování vybraných metrik služeb naprogramovaných v jazyku Go
//    https://www.root.cz/clanky/sledovani-vybranych-metrik-sluzeb-naprogramovanych-v-jazyku-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze třicáté první části:
//    https://github.com/tisnik/go-root/blob/master/article_31/README.md
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
