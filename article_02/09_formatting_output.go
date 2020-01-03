// Seriál "Programovací jazyk Go"
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 9:
//    Základní formátování celočíselných hodnot na výstupu

package main

import "fmt"

func main() {
	var a uint8 = 20
	var b uint16 = 2000
	var c uint32 = 20000
	var d uint32 = 2000000

	fmt.Println("%d")
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", c)
	fmt.Printf("%d\n", d)

	fmt.Println("\n%5d")
	fmt.Printf("%5d\n", a)
	fmt.Printf("%5d\n", b)
	fmt.Printf("%5d\n", c)
	fmt.Printf("%5d\n", d)

	fmt.Println("\n%05d")
	fmt.Printf("%05d\n", a)
	fmt.Printf("%05d\n", b)
	fmt.Printf("%05d\n", c)
	fmt.Printf("%05d\n", d)

	fmt.Println("\n%-5d")
	fmt.Printf("%-5d\n", a)
	fmt.Printf("%-5d\n", b)
	fmt.Printf("%-5d\n", c)
	fmt.Printf("%-5d\n", d)

	fmt.Println("\n%+5d")
	fmt.Printf("%+5d\n", a)
	fmt.Printf("%+5d\n", b)
	fmt.Printf("%+5d\n", c)
	fmt.Printf("%+5d\n", d)

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
