// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtvrtá část
//    Rozhraní, metody, gorutiny a kanály v programovacím jazyku Go
//    https://www.root.cz/clanky/rozhrani-metody-gorutiny-a-kanaly-v-programovacim-jazyku-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_04/README.md
//
// Demonstrační příklad číslo 16:
//    Zavolání další gorutiny z jiné gorutiny.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_04/16_goroutine_from_goroutine.html

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
	go printChars()
	go printDots()
	for i := 0; i < 60; i++ {
		fmt.Print(" ")
		time.Sleep(110 * time.Millisecond)
	}
}

func main() {
	fmt.Println("main begin")
	go printSpaces()
	time.Sleep(6 * time.Second)
	fmt.Println("main end")
}
