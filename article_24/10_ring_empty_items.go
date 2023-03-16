// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá čtvrtá část
//     Kontejnery v základní knihovně programovacího jazyka Go
//     https://www.root.cz/clanky/kontejnery-v-zakladni-knihovne-programovaciho-jazyka-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z dvacáté čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_24/README.md
//
// Demonstrační příklad číslo 10:
//     	Cyklická fronta s některými prázdnými prvky

package main

import (
	"container/ring"
	"fmt"
)

func printRing(r *ring.Ring) {
	length := r.Len()
	for i := 0; i < length*2; i++ {
		fmt.Println(i, r.Value)
		r = r.Next()
	}
}

func main() {
	r := ring.New(10)
	r.Value = "foo"
	r = r.Next()
	r.Value = "bar"
	r = r.Next()
	r.Value = "baz"
	printRing(r)
}
