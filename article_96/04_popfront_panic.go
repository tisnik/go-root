package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

func main() {
	var q deque.Deque[int]

	for i := 1; i <= 10; i++ {
		q.PushBack(100 + i)
		fmt.Println("Deque length:  ", q.Len())
		fmt.Println("Deque capacity:", q.Cap())
		fmt.Println("Deque value:   ", q)
		fmt.Println()
	}

	fmt.Println()

	for true {
		fmt.Println("Pop value:     ", q.PopFront())
		fmt.Println("Deque length:  ", q.Len())
		fmt.Println("Deque capacity:", q.Cap())
		fmt.Println("Deque value:   ", q)
		fmt.Println()
	}
}
