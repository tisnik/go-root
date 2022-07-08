// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/

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

	it := heap.Iterator()
	for it.Next() {
		value := it.Value()
		fmt.Printf("%3d \t %T\n", value, value)
	}
}
