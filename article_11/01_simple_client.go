// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 1:
//    Klient, který přečte ze serveru sekvenci bajtů
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_11/01_simple_client.html

package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		println("Connection refused!")
	} else {
		var b [1]byte
		n, err := conn.Read(b[:])
		if err != nil {
			println("No response!")
		} else {
			if n == 1 {
				fmt.Printf("Received %d byte: %v\n", n, b)
			} else {
				fmt.Printf("Received %d bytes: %v\n", n, b)
			}
		}
	}
}
