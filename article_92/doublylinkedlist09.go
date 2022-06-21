package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/doublylinkedlist"
)

func main() {
	l := doublylinkedlist.New[string]("foo", "bar", "baz")
	fmt.Println(l)
	fmt.Println()

	l.Swap(0, 1)
	fmt.Println(l)
	fmt.Println()

	l.Swap(1, 2)
	fmt.Println(l)
	fmt.Println()

	l.Swap(2, 3)
	fmt.Println(l)
	fmt.Println()
}
