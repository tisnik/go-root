// Seriál "Programovací jazyk Go"
//
// Druhá část
//
// Demonstrační příklad číslo 8:
//    Základní formátování celočíselných hodnot na výstupu

package main

import "fmt"

func main() {
	var a int8 = -10
	var b int16 = -1000
	var c int32 = -10000
	var d int32 = -1000000

        fmt.Println("%d")
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", c)
	fmt.Printf("%d\n", d)

        fmt.Println("\n%x")
	fmt.Printf("%x\n", a)
	fmt.Printf("%x\n", b)
	fmt.Printf("%x\n", c)
	fmt.Printf("%x\n", d)

        fmt.Println("\n%X")
	fmt.Printf("%X\n", a)
	fmt.Printf("%X\n", b)
	fmt.Printf("%X\n", c)
	fmt.Printf("%X\n", d)

        fmt.Println("\n%b")
	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", b)
	fmt.Printf("%b\n", c)
	fmt.Printf("%b\n", d)
}
