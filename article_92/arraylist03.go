package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/arraylist"
)

func main() {
	l := arraylist.New[string](1, 2, 3)
	fmt.Println(l)
}
