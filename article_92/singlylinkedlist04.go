package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/singlylinkedlist"
)

func main() {
	l := singlylinkedlist.New[string]("a", "b", "c")
	fmt.Println(l)

	l.Add(1)
	fmt.Println(l)

	l.Add(2, 3)
	fmt.Println(l)
}
