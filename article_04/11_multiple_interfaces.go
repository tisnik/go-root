// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtvrtá část
//    Rozhraní, metody, gorutiny a kanály v programovacím jazyku Go
//    https://www.root.cz/clanky/rozhrani-metody-gorutiny-a-kanaly-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 11:
//    Typ implementující dvě rozhraní.

package main

import "fmt"

// Interface1 je exportované rozhraní s jedinou metodou
type Interface1 interface {
	method1()
}

// Interface2 je exportované rozhraní s jedinou metodou
type Interface2 interface {
	method2()
}

// Type je uživatelsky definovaný datový typ
type Type struct{}

func (Type) method1() {
	fmt.Println("Type.method1")
}

func (Type) method2() {
	fmt.Println("Type.method2")
}

func f1(i Interface1) {
	fmt.Println("Interface1.f1")
	i.method1()
}

func f2(i Interface2) {
	fmt.Println("Interface2.f2")
	i.method2()
}

func main() {
	t := Type{}

	t.method1()
	t.method2()
	fmt.Println()

	f1(t)
	fmt.Println()

	f2(t)
	fmt.Println()
}
