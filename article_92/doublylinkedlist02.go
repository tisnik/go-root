package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/doublylinkedlist"
)

func main() {
	l := doublylinkedlist.New[string]("a", "b", "c")
	fmt.Println(l)

	l.Add("a")
	fmt.Println(l)

	l.Add("b", "c")
	fmt.Println(l)
}
