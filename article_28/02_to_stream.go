// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá osmá část
//    Datové struktury s líným vyhodnocováním v programovacím jazyce Go
//    https://www.root.cz/clanky/datove-struktury-s-linym-vyhodnocovanim-v-programovacim-jazyce-go/
//
// Seznam příkladů z dvacáté osmé části:
//    https://github.com/tisnik/go-root/blob/master/article_28/README.md
//
// Demonstrační příklad číslo 2:
//    Konverze polí a řezů do streamů.

package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func main() {
	values := []int{1, 2, 3, 4, 5}
	fmt.Printf("input:    %v\n", values)

	stream := koazee.StreamOf(values)
	fmt.Printf("stream:   %v\n", stream.Out().Val())

	fmt.Println()

	words := []string{"one", "two", "three", "four", "five"}
	fmt.Printf("words:    %v\n", words)

	stream = koazee.StreamOf(words)
	fmt.Printf("stream:   %v\n", stream.Out().Val())

	fmt.Println()

	anything := []interface{}{"one", 1, 1.0, 1 + 1i}
	fmt.Printf("anything: %v\n", anything)

	stream3 := koazee.StreamOf(anything)
	fmt.Printf("stream:   %v\n", stream3.Out().Val())
}
