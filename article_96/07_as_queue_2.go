package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

func main() {
	var q deque.Deque[int]

	for i := 1; i <= 10; i++ {
		q.PushFront(100 + i*2)
		q.PushFront(101 + i*2)
		fmt.Println(q.PopBack())
	}

	for q.Len() > 0 {
		fmt.Println(q.PopBack())
	}
}
