// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devadesátá třetí část
//    Knihovny s implementací generických datových typů pro programovací jazyk Go (2)
//    https://www.root.cz/clanky/knihovny-s-implementaci-generickych-datovych-typu-pro-programovaci-jazyk-go-2/
//
// Seznam příkladů ze devadesáté třetí části:
//    https://github.com/tisnik/go-root/blob/master/article_93/README.md

package main

import (
	"fmt"

	heapImpl "github.com/daichi-m/go18ds/trees/binaryheap"
)

func main() {
	heap := heapImpl.NewWithIntComparator()
	fmt.Println(heap)

	heap.Push(1)
	fmt.Println(heap)

	heap.Push(2)
	heap.Push(3)
	fmt.Println(heap)

	heap.Push(4)
	heap.Push(5)
	fmt.Println(heap)

	heap.Push(6)
	heap.Push(7)
	fmt.Println(heap)

	heap.Push(8)
	heap.Push(9)
	fmt.Println(heap)
}
