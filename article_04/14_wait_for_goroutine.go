// Seriál "Programovací jazyk Go"
//
// Čtvrtá část
//
// Demonstrační příklad číslo 14:
//    Zavolání funkce v gorutině, počkání na dokončení.

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
