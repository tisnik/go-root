// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtvrtá část
//    Rozhraní, metody, gorutiny a kanály v programovacím jazyku Go
//    https://www.root.cz/clanky/rozhrani-metody-gorutiny-a-kanaly-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 14:
//    Zavolání funkce v gorutině, počkání na dokončení.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_04/14_wait_for_goroutine.html

package main

import (
	"fmt"
	"time"
)

func message(id int) {
	fmt.Printf("gorutina %d\n", id)
}

func main() {
	fmt.Println("main begin")
	go message(1)
	go message(2)
	time.Sleep(2 * time.Second)
	fmt.Println("main end")
}
