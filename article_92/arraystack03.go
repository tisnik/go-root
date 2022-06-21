package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/stacks/arraystack"
)

func main() {
	stack := arraystack.New[string]()
	stack.Push("foo")
	stack.Push("bar")
	stack.Push("baz")

	for i := 0; i < 10; i++ {
		value, exists := stack.Peek()
		fmt.Println(i, value, exists)
		stack.Pop()
	}
}