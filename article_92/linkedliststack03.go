// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devadesátá druhá část
//    Knihovny s implementací generických datových typů pro programovací jazyk Go (2)
//    https://www.root.cz/clanky/knihovny-s-implementaci-generickych-datovych-typu-pro-programovaci-jazyk-go-2/
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

	for i := 0; i < 10; i++ {
		value, exists := stack.Peek()
		fmt.Println(i, value, exists)
		stack.Pop()
	}
}
