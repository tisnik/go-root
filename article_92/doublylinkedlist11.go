package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/doublylinkedlist"
)

func main() {
	l := doublylinkedlist.New[string]("zoo", "aleph", "foo", "bar", "baz")

	it := l.Iterator()
	for it.Next() {
		value := it.Value()
		fmt.Printf("%3s \t %T\n", value, value)
	}

}
