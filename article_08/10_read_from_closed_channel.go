// Seriál "Programovací jazyk Go"
//
// Osmá část
//     Tvorba balíčků a pokročilejší operace s kanály v jazyce Go
//     https://www.root.cz/clanky/tvorba-balicku-a-pokrocilejsi-operace-s-kanaly-v-jazyce-go/
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
