// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtvrtá část
//    Rozhraní, metody, gorutiny a kanály v programovacím jazyku Go
//    https://www.root.cz/clanky/rozhrani-metody-gorutiny-a-kanaly-v-programovacim-jazyku-go/
//
// Seznam příkladů ze čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_04/README.md
//
// Demonstrační příklad číslo 12:
//    Typ implementující dvě rozhraní.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_04/12_embedded_interface.html

package main

import "fmt"

// Interface1 je exportované rozhraní s jedinou metodou
type Interface1 interface {
	method1()
}

// Interface2 je exportované rozhraní s jedinou metodou
type Interface2 interface {
	Interface1
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
