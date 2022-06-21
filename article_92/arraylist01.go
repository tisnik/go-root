package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/arraylist"
)

func main() {
	l := arraylist.New[string]("a", "b", "c")
	fmt.Println(l)
}
