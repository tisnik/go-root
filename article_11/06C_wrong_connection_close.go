// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 6C:
//     Připojení je nutné ukončit v gorutině, ne mimo ni
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_11/06C_wrong_connection_close.html

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
		defer conn.Close()
		go func(c net.Conn) {
			fmt.Fprintf(c, "Holla\n")
		}(conn)
	}
}
