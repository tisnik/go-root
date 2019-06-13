// Seriál "Programovací jazyk Go"
//
// Dvacátá čtvrtá část
//
// Demonstrační příklad číslo 11:
//     	Iterace nad prvky cyklické fronty

package main

import (
	"container/ring"
	"fmt"
)

func printItem(item interface{}) {
	fmt.Println(item)
}

func main() {
	r := ring.New(3)
	r.Value = "foo"
	r = r.Next()
	r.Value = "bar"
	r = r.Next()
	r.Value = "baz"
	r.Do(printItem)
}
