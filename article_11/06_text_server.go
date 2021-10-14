// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 6:
//     Server posílající klientovi textová data
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_11/06_text_server.html

package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		println("Can't open the port!")
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			println("Connection refused!")
		}
		go func(c net.Conn) {
			fmt.Fprintf(c, "Holla\n")
			c.Close()
		}(conn)
	}
}
