// Seriál "Programovací jazyk Go"
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 9:
//     Konstruktor adresy typu IPv4

package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Printf("%v\n", net.IPv4(127, 0, 0, 1))
	fmt.Printf("%v\n", net.IPv4(192, 168, 10, 3))
}
