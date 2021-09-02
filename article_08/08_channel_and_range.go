// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmá část
//     Tvorba balíčků a pokročilejší operace s kanály v jazyce Go
//     https://www.root.cz/clanky/tvorba-balicku-a-pokrocilejsi-operace-s-kanaly-v-jazyce-go/
//
// Demonstrační příklad číslo 8:
//    Kanál s definovanou kapacitou, příkaz range.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_08/08_channel_and_range.html

package main

import (
	"fmt"
)

func fillInChannel(channel chan rune) {
	channel <- 'A'
	channel <- 'B'
	channel <- 'C'
}

func main() {
	channel := make(chan rune, 3)

	go fillInChannel(channel)

	for msg := range channel {
		fmt.Printf("%c ", msg)
	}
}
