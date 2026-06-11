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
// Demonstrační příklad číslo 12:
//    Balíček container/ring.
//    Vylepšená iterace nad prvky cyklické fronty.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_24/12_ring_do_closure.html

package main

import (
	"container/ring"
	"fmt"
)

func printItem(index int, item interface{}) {
	fmt.Println(index, item)
}

func main() {
	r := ring.New(3)
	r.Value = "foo"
	r = r.Next()
	r.Value = "bar"
	r = r.Next()
	r.Value = "baz"

	i := 1
	r.Do(func(item interface{}) {
		printItem(i, item)
		i++
	})
}
