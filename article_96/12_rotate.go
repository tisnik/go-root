package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

func main() {
	var q deque.Deque[int]

	for i := 1; i <= 10; i++ {
		q.PushBack(100 + i)
	}

	q.Rotate(-3)

	for q.Len() > 0 {
		fmt.Println(q.PopFront())
	}
}
