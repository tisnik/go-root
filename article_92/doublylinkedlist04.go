package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/doublylinkedlist"
)

func main() {
	l := doublylinkedlist.New[string]("a", "b", "c")
	fmt.Println(l)

	l.Add(1)
	fmt.Println(l)

	l.Add(2, 3)
	fmt.Println(l)
}
