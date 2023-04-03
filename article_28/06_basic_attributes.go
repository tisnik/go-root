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
// Demonstrační příklad číslo 6:
//    Základní atributy streamů.

package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func main() {
	values := []int{1, 2, 3, 4, 5}
	fmt.Printf("input:\t%v\n", values)

	stream := koazee.StreamOf(values)

	cnt, _ := stream.Count()
	fmt.Printf("count:\t%d\n", cnt)

	first := stream.First().Val()
	fmt.Printf("first:\t%d\n", first)

	second := stream.At(1).Val()
	fmt.Printf("second:\t%d\n", second)

	last := stream.Last().Val()
	fmt.Printf("last:\t%d\n", last)

	unknown := stream.At(10).Val()
	fmt.Printf("unknown:\t%v\n", unknown)
}
