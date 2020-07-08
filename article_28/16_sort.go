// Seriál "Programovací jazyk Go"
//
// Dvacátá osmá část
//    Datové struktury s líným vyhodnocováním v programovacím jazyce Go
//    https://www.root.cz/clanky/datove-struktury-s-linym-vyhodnocovanim-v-programovacim-jazyce-go/
//
// Demonstrační příklad číslo 16:
//    Metoda Sort().

package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func compare1(x, y int) int {
	return x - y
}

func compare2(x, y int) int {
	return y - x
}

func main() {
	values := []int{100, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -1}
	fmt.Printf("input: %v\n", values)

	stream1 := koazee.StreamOf(values)
	fmt.Printf("stream1:  %v\n", stream1.Out().Val())

	stream1.Sort(compare1).Do()
	fmt.Printf("stream1:  %v\n", stream1.Out().Val())

	stream1.Sort(compare2).Do()
	fmt.Printf("stream1:  %v\n", stream1.Out().Val())
}
