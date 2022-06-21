package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/singlylinkedlist"
)

func testGet(l *singlylinkedlist.List[string]) {
	fmt.Println("List:", l)

	for i := -1; i < l.Size()+1; i++ {
		value, found := l.Get(i)
		if found {
			fmt.Printf("Item %d = '%s'\n", i, value)
		} else {
			fmt.Printf("Item %d not found\n", i)
		}
	}
	fmt.Println()
}

func main() {
	l := singlylinkedlist.New[string]()
	testGet(l)

	l.Add("a")
	testGet(l)

	l.Add("b", "c")
	testGet(l)

	l.Clear()
	testGet(l)
}
