package main

import (
	"fmt"

	heapImpl "github.com/daichi-m/go18ds/trees/binaryheap"
)

func main() {
	heap := heapImpl.NewWithIntComparator()
	fmt.Println(heap)

	heap.Push(9)
	heap.Push(8)
	fmt.Println(heap)

	heap.Push(7)
	heap.Push(6)
	fmt.Println(heap)

	heap.Push(5)
	heap.Push(4)
	fmt.Println(heap)

	heap.Push(3)
	heap.Push(2)
	fmt.Println(heap)

	heap.Push(1)
	fmt.Println(heap)
}