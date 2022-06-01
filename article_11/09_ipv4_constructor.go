// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
//
// Seznam příkladů z jedenácté části:
//    https://github.com/tisnik/go-root/blob/master/article_11/README.md
//
// Demonstrační příklad číslo 9:
//     Konstruktor adresy typu IPv4
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_11/09_ipv4_constructor.html

package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Printf("%v\n", net.IPv4(127, 0, 0, 1))
	fmt.Printf("%v\n", net.IPv4(192, 168, 10, 3))
}
