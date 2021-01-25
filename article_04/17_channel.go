// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtvrtá část
//    Rozhraní, metody, gorutiny a kanály v programovacím jazyku Go
//    https://www.root.cz/clanky/rozhrani-metody-gorutiny-a-kanaly-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 17:
//    Kanál pro komunikaci a synchronizaci mezi gorutinami.

package main

import "fmt"

func message(id int, channel chan int) {
	fmt.Printf("gorutina %d\n", id)
	channel <- 1
}

func main() {
	channel := make(chan int)

	fmt.Println("main begin")
	go message(1, channel)

	fmt.Println("waiting...")

	code, status := <-channel

	fmt.Printf("received code: %d and status: %t\n", code, status)
	fmt.Println("main end")
}
