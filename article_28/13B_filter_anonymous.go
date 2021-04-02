// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá osmá část
//    Datové struktury s líným vyhodnocováním v programovacím jazyce Go
//    https://www.root.cz/clanky/datove-struktury-s-linym-vyhodnocovanim-v-programovacim-jazyce-go/
//
// Demonstrační příklad číslo 13B:
//    Metoda Filter() a stream s celočíselnými hodnotami.

package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("input: %v\n", values)

	stream1 := koazee.StreamOf(values)
	stream2 := stream1.Filter(func(x int) bool { return x%2 == 0 }).Do()
	stream3 := stream1.Filter(func(x int) bool { return x%3 == 0 }).Do()

	fmt.Printf("stream1:  %v\n", stream1.Out().Val())
	fmt.Printf("stream2:  %v\n", stream2.Out().Val())
	fmt.Printf("stream3:  %v\n", stream3.Out().Val())
}
