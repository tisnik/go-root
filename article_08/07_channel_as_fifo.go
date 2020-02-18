// Seriál "Programovací jazyk Go"
//
// Osmá část
//     Tvorba balíčků a pokročilejší operace s kanály v jazyce Go
//     https://www.root.cz/clanky/tvorba-balicku-a-pokrocilejsi-operace-s-kanaly-v-jazyce-go/
//
// Demonstrační příklad číslo 7:
//    Kanál s definovanou kapacitou, který se chová jako FIFO.

package main

import (
	"fmt"
)

func main() {
	channel := make(chan rune, 3)
	channel <- 'A'
	channel <- 'B'
	channel <- 'C'

	for i := 0; i < 3; i++ {
		fmt.Printf("%c ", <-channel)
	}
}
