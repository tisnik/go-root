package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/doublylinkedlist"
)

func testContains(l *doublylinkedlist.List[string]) {
	fmt.Println("List:", l)
	for _, c := range []string{"a", "b", "c", "x"} {
		fmt.Printf("Contains '%s': %t\n", c, l.Contains(c))
	}
	fmt.Println()
}

func main() {
	l := doublylinkedlist.New[string]()
	testContains(l)

	l.Add("a")
	testContains(l)

	l.Add("b", "c")
	testContains(l)

	l.Clear()
	testContains(l)
}
