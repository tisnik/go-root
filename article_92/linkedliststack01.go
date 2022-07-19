// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z devadesáté druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_92/README.md

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
