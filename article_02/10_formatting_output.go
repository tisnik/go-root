// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 10:
//    Základní formátování celočíselných hodnot na výstupu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_02/10_formatting_output.html

package main

import "fmt"

func main() {
	var a int8 = -10
	var b int16 = -1000
	var c int32 = -10000
	var d int32 = -1000000

	var r1 rune = 'a'
	var r2 rune = '\x40'
	var r3 rune = '\n'
	var r4 rune = '\u03BB'

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

	fmt.Println("%c")
	fmt.Printf("%c\n", r1)
	fmt.Printf("%c\n", r2)
	fmt.Printf("%c\n", r3)
	fmt.Printf("%c\n", r4)
}
