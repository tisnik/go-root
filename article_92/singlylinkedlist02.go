package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/singlylinkedlist"
)

func main() {
	l := singlylinkedlist.New[string]("a", "b", "c")
	fmt.Println(l)

	l.Add("a")
	fmt.Println(l)

	l.Add("b", "c")
	fmt.Println(l)
}
