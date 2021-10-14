// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 7:
//     Překlad doménového jména na IP adresy
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_11/07_lookup.html

package main

import (
	"fmt"
	"net"
)

func performLookup(address string) {
	ip, err := net.LookupIP(address)
	if err != nil {
		println("Lookup failure")
	} else {
		fmt.Printf("%v\n", ip)
	}
}

func main() {
	performLookup("localhost")
	performLookup("root.cz")
	performLookup("google.com")
}
