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
// Demonstrační příklad číslo 13:
//    Zavolání funkce v gorutině.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_04/13_simple_goroutine.html

package main

import "fmt"

func message(id int) {
	fmt.Printf("gorutina %d\n", id)
}

func main() {
	fmt.Println("main begin")
	go message(1)
	go message(2)
	fmt.Println("main end")
}
