package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

func main() {
	var q deque.Deque[int]

	for i := 1; i <= 10; i++ {
		q.PushBack(100 + i*2)
		q.PushBack(101 + i*2)
		fmt.Println(q.PopFront())
	}

	for q.Len() > 0 {
		fmt.Println(q.PopFront())
	}
}
