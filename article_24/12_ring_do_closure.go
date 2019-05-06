package main

import (
	"container/ring"
	"fmt"
)

func printItem(index int, item interface{}) {
	fmt.Println(index, item)
}

func main() {
	r := ring.New(3)
	r.Value = "foo"
	r = r.Next()
	r.Value = "bar"
	r = r.Next()
	r.Value = "baz"

	i := 1
	r.Do(func(item interface{}) {
		printItem(i, item)
		i++
	})
}
