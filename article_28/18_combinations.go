// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá osmá část
//    Datové struktury s líným vyhodnocováním v programovacím jazyce Go
//    https://www.root.cz/clanky/datove-struktury-s-linym-vyhodnocovanim-v-programovacim-jazyce-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z dvacáté osmé části:
//    https://github.com/tisnik/go-root/blob/master/article_28/README.md
//
// Demonstrační příklad číslo 18:
//    Kombinace různých metod pro "streaming".

package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func printInt(x int) {
	fmt.Printf("%d\n", x)
}

func doubleValue(x int) int {
	return x * 2
}

func negate(x int) int {
	return -x
}

func evenValue(x int) bool {
	return x%2 == 0
}

func divisibleBy3(x int) bool {
	return x%3 == 0
}

func main() {
	values1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Printf("input #1: %v\n", values1)

	stream1 := koazee.StreamOf(values1).
		Filter(evenValue).
		Filter(divisibleBy3).
		Map(negate).
		Map(doubleValue)
	stream1.ForEach(printInt).Do()
}
