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

	for {
		value, exists := stack.Pop()
		if exists {
			fmt.Println(value, stack.Size(), stack.Empty())
		} else {
			break
		}
	}
}
