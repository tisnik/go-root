package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/stacks/arraystack"
)

func main() {
	stack := arraystack.New[string]()
	fmt.Println(stack)

	stack.Push("foo")
	fmt.Println(stack)

	stack.Push("bar")
	fmt.Println(stack)

	stack.Push("baz")
	fmt.Println(stack)
}
