// Seriál "Programovací jazyk Go"
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 7:
//     Překlad doménového jména na IP adresy

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
