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
// Demonstrační příklad číslo 20:
//     	Úprava předchozího příkladu, aby zobrazoval prováděné operace

package main

import (
	"container/heap"
	"fmt"
)

// StringHeap represents simple heap abstract data structure based on array
type StringHeap []string

func (h StringHeap) Len() int {
	return len(h)
}

func (h StringHeap) Less(i, j int) bool {
	fmt.Printf("compare %s < %s\n", h[i], h[j])
	return h[i] < h[j]
}

func (h StringHeap) Swap(i, j int) {
	fmt.Printf("swap    %s <-> %s\n", h[i], h[j])
	h[i], h[j] = h[j], h[i]
}

func (h *StringHeap) Push(x interface{}) {
	*h = append(*h, x.(string))
}

func (h *StringHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &StringHeap{}

	heap.Init(h)
	heap.Push(h, "a")
	heap.Push(h, "z")
	heap.Push(h, "c")
	heap.Push(h, "b")
	heap.Push(h, "d")
	heap.Push(h, "x")

	fmt.Println("\n-----------------------------")
	fmt.Printf("First item: %s\n", (*h)[0])
	i := 0
	for h.Len() > 0 {
		i++
		fmt.Printf("item #%d = %s\n", i, heap.Pop(h))
	}
}
