// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtvrtá část
//    Rozhraní, metody, gorutiny a kanály v programovacím jazyku Go
//    https://www.root.cz/clanky/rozhrani-metody-gorutiny-a-kanaly-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 15:
//    Zavolání funkcí v gorutinách, počkání na dokončení.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_04/15_multiple_goroutines.html

package main

import (
	"fmt"
	"time"
)

func printChars() {
	for ch := 'a'; ch <= 'z'; ch++ {
		fmt.Printf("%c", ch)
		time.Sleep(200 * time.Millisecond)
	}
}

func printDots() {
	for i := 0; i < 30; i++ {
		fmt.Print(".")
		time.Sleep(200 * time.Millisecond)
	}
}

func printSpaces() {
	for i := 0; i < 60; i++ {
		fmt.Print(" ")
		time.Sleep(110 * time.Millisecond)
	}
}

func main() {
	fmt.Println("main begin")
	go printChars()
	go printSpaces()
	go printDots()
	time.Sleep(6 * time.Second)
	fmt.Println("main end")
}
