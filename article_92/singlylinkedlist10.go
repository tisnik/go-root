package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/singlylinkedlist"
	"github.com/daichi-m/go18ds/utils"
)

func main() {
	l := singlylinkedlist.New[string]("zoo", "aleph", "foo", "bar", "baz")
	fmt.Println(l)
	fmt.Println()

	l.Sort(utils.StringComparator)
	fmt.Println(l)
}
