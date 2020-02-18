// Seriál "Programovací jazyk Go"
//
// Osmá část
//     Tvorba balíčků a pokročilejší operace s kanály v jazyce Go
//     https://www.root.cz/clanky/tvorba-balicku-a-pokrocilejsi-operace-s-kanaly-v-jazyce-go/
//
// Demonstrační příklad číslo 9:
//    Kanál s definovanou kapacitou, příkaz range.

package main

import (
	"fmt"
)

func fill_in_channel(channel chan rune) {
	channel <- 'A'
	channel <- 'B'
	channel <- 'C'
	close(channel)
}

func main() {
	channel := make(chan rune, 3)

	go fill_in_channel(channel)

	for msg := range channel {
		fmt.Printf("%c ", msg)
	}
}
