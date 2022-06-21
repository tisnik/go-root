package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/doublylinkedlist"
	"github.com/daichi-m/go18ds/utils"
)

func main() {
	l := doublylinkedlist.New[string]("zoo", "aleph", "foo", "bar", "baz")
	fmt.Println(l)
	fmt.Println()

	l.Sort(utils.StringComparator)
	fmt.Println(l)
}
