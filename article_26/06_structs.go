// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//
// Demonstrační příklad číslo 6:
//    Struktury (záznamy)

package main

import "fmt"

func main() {
	var s struct {
		a int
		b bool
		c chan int
		d []int
	}

	fmt.Println(s)
}
