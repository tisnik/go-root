// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmá část
//     Tvorba balíčků a pokročilejší operace s kanály v jazyce Go
//     https://www.root.cz/clanky/tvorba-balicku-a-pokrocilejsi-operace-s-kanaly-v-jazyce-go/
//
// Demonstrační příklad číslo 10:
//    Čtení ze zavřeného kanálu.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_08/10_read_from_closed_channel.html

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
