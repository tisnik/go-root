// Seriál "Programovací jazyk Go"
//
// Čtvrtá část
//
// Demonstrační příklad číslo 16:
//    Zavolání další gorutiny z jiné gorutiny.

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
	go print_chars()
	go print_dots()
	for i := 0; i < 60; i++ {
		fmt.Print(" ")
		time.Sleep(110 * time.Millisecond)
	}
}

func main() {
	fmt.Println("main begin")
	go print_spaces()
	time.Sleep(6 * time.Second)
	fmt.Println("main end")
}
