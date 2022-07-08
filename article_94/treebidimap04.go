// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/

package main

import (
	"fmt"

	mapImpl "github.com/daichi-m/go18ds/maps/treebidimap"
	"github.com/daichi-m/go18ds/utils"
)

func main() {
	m := mapImpl.NewWith(utils.NumberComparator[int], utils.NumberComparator[int])

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
