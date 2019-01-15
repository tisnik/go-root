// Seriál "Programovací jazyk Go"
//
// Osmá část
//
// Demonstrační příklad číslo 10:
//    Čtení ze zavřeného kanálu.

package main

import (
	"fmt"
)

func main() {
	channel := make(chan int, 3)
	close(channel)

	fmt.Printf("%d\n", <-channel)
	fmt.Printf("%d\n", <-channel)
	fmt.Printf("%d\n", <-channel)
}
