package main

import (
	"fmt"

	mapImpl "github.com/daichi-m/go18ds/maps/hashbidimap"
)

func main() {
	m := mapImpl.New[int, int]()

	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)

	fmt.Println("Size", m.Size())
	fmt.Println("Keys", m.Keys())
	fmt.Println("Values", m.Values())

	fmt.Println()

	m.Put(4, 1)

	fmt.Println("Size", m.Size())
	fmt.Println("Keys", m.Keys())
	fmt.Println("Values", m.Values())
}
