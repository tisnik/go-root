// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmá část
//     Tvorba balíčků a pokročilejší operace s kanály v jazyce Go
//     https://www.root.cz/clanky/tvorba-balicku-a-pokrocilejsi-operace-s-kanaly-v-jazyce-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z osmé části:
//    https://github.com/tisnik/go-root/blob/master/article_08/README.md
//
// Demonstrační příklad číslo 7:
//    Kanál s definovanou kapacitou, který se chová jako FIFO.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_08/07_channel_as_fifo.html

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
