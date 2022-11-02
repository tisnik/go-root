package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

func main() {
	var q deque.Deque[int]

	fmt.Println("Deque length:  ", q.Len())
	fmt.Println("Deque capacity:", q.Cap())
	fmt.Println("Deque value:   ", q)
}
