// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá osmá část
//    Datové struktury s líným vyhodnocováním v programovacím jazyce Go
//    https://www.root.cz/clanky/datove-struktury-s-linym-vyhodnocovanim-v-programovacim-jazyce-go/
//
// Demonstrační příklad číslo 20:
//    Metoda Reduce().

package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func add(x, y int) int {
	return x + y
}

func main() {
	values1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("input #1: %v\n", values1)

	sum := koazee.StreamOf(values1).Reduce(add)
	fmt.Printf("sum: %d\n", sum.Val())
}
