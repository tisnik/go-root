package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/stacks/linkedliststack"
)

func main() {
	stack := linkedliststack.New[string]()
	fmt.Println(stack)

	stack.Push("foo")
	fmt.Println(stack)

	stack.Push("bar")
	fmt.Println(stack)

	stack.Push("baz")
	fmt.Println(stack)
}
