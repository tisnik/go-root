// Seriál "Programovací jazyk Go"
//
// Čtvrtá část
//
// Demonstrační příklad číslo 13:
//    Zavolání funkce v gorutině.

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
