package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

func main() {
	var q deque.Deque[int]

	fmt.Println("Length Capacity")
	for i := 1; i <= 1000; i++ {
		q.PushBack(i)
		fmt.Printf("%5d %5d\n", q.Len(), q.Cap())
	}
}
