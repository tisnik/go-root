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
	stack.Push("foo")
	stack.Push("bar")
	stack.Push("baz")

	for {
		value, exists := stack.Pop()
		if exists {
			fmt.Println(value)
		} else {
			break
		}
	}
}
