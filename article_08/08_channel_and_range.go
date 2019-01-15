// Seriál "Programovací jazyk Go"
//
// Osmá část
//
// Demonstrační příklad číslo 8:
//    Kanál s definovanou kapacitou, příkaz range.

package main

import (
	"fmt"
)

func fill_in_channel(channel chan rune) {
	channel <- 'A'
	channel <- 'B'
	channel <- 'C'
}

func main() {
	channel := make(chan rune, 3)

	go fill_in_channel(channel)

	for msg := range channel {
		fmt.Printf("%c ", msg)
	}
}
