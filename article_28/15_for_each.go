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
// Demonstrační příklad číslo 15:
//    Metoda ForEach().

package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func printInt(x int) {
	fmt.Printf("%d\t", x)
}

func main() {
	values1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("input #1: %v\n", values1)

	stream1 := koazee.StreamOf(values1)
	stream1.ForEach(printInt)

	fmt.Println()

	values2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("input #2: %v\n", values2)

	stream2 := koazee.StreamOf(values2)
	stream2.ForEach(printInt).Do()
}
