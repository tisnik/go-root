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
// Demonstrační příklad číslo 8:
//    Balíček container/ring.
//    Prázdná cyklická fronta.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_24/08_empty_ring.html

package main

import (
	"container/ring"
	"fmt"
)

func printRing(r *ring.Ring) {
	length := r.Len()
	for i := 0; i < length; i++ {
		fmt.Println(i, r.Value)
		r = r.Next()
	}
}

func main() {
	r := ring.New(10)
	printRing(r)
}
