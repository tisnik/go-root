package main

import (
	"container/heap"
	"fmt"
)

type StringHeap []string

func (h StringHeap) Len() int {
	return len(h)
}

func (h StringHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h StringHeap) Swap(i, j int) {
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
	h := &StringHeap{"foo", "bar", "baz", "zzz", "aaa"}
	heap.Init(h)
	heap.Push(h, "ZZZ")
	heap.Push(h, "AAA")
	fmt.Printf("First item: %s\n", (*h)[0])
	i := 0
	for h.Len() > 0 {
		i++
		fmt.Printf("item #%d = %s\n", i, heap.Pop(h))
	}
}
