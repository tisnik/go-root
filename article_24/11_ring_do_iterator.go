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
// Demonstrační příklad číslo 11:
//    Balíček container/ring.
//    Iterace nad prvky cyklické fronty.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_24/11_ring_do_iterator.html

package main

import (
	"container/ring"
	"fmt"
)

func printItem(item interface{}) {
	fmt.Println(item)
}

func main() {
	r := ring.New(3)
	r.Value = "foo"
	r = r.Next()
	r.Value = "bar"
	r = r.Next()
	r.Value = "baz"
	r.Do(printItem)
}
