// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá pátá část
//     Ladění aplikací v Go s využitím GNU Debuggeru a debuggeru Delve
//     https://www.root.cz/clanky/ladeni-aplikaci-v-go-s-vyuzitim-gnu-debuggeru-a-debuggeru-delve/
//
// Demonstrační příklad číslo 6:
//     Gorutiny a kanály

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
