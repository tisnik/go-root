// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//    Problematika nulových hodnot v Go, aneb proč nil != nil
//    https://www.root.cz/clanky/problematika-nulovych-hodnot-v-go-aneb-proc-nil-nil/
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
