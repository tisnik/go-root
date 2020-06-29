// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//    Problematika nulových hodnot v Go, aneb proč nil != nil
//    https://www.root.cz/clanky/problematika-nulovych-hodnot-v-go-aneb-proc-nil-nil/
//
// Demonstrační příklad číslo 5:
//    Ostatní datové typy, lepší přístup

package main

import "fmt"

func main() {
	var a [10]complex64
	var p *int
	var s []int
	var m map[string]int
	var c chan int
	var f func()
	var i interface{}

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(s)
	fmt.Println(m)
	fmt.Println(c)
	fmt.Println(f)
	fmt.Println(i)
}
