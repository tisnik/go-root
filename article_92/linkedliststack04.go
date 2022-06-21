package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/stacks/linkedliststack"
)

func main() {
	stack := linkedliststack.New[string]()
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
