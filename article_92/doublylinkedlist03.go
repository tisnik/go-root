package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/doublylinkedlist"
)

func main() {
	l := doublylinkedlist.New[string](1, 2, 3)
	fmt.Println(l)
}
