// Seriál "Programovací jazyk Go"
//
// Čtvrtá část
//    Rozhraní, metody, gorutiny a kanály v programovacím jazyku Go
//    https://www.root.cz/clanky/rozhrani-metody-gorutiny-a-kanaly-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 15:
//    Zavolání funkcí v gorutinách, počkání na dokončení.

package main

import (
	"fmt"
	"time"
)

func print_chars() {
	for ch := 'a'; ch <= 'z'; ch++ {
		fmt.Printf("%c", ch)
		time.Sleep(200 * time.Millisecond)
	}
}

func print_dots() {
	for i := 0; i < 30; i++ {
		fmt.Print(".")
		time.Sleep(200 * time.Millisecond)
	}
}

func print_spaces() {
	for i := 0; i < 60; i++ {
		fmt.Print(" ")
		time.Sleep(110 * time.Millisecond)
	}
}

func main() {
	fmt.Println("main begin")
	go print_chars()
	go print_spaces()
	go print_dots()
	time.Sleep(6 * time.Second)
	fmt.Println("main end")
}
