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
// Demonstrační příklad číslo 2B:
//     Úprava adresy v předchozím příkladu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_11/02B_simple_server_no_localhost.html

package main

import (
	"net"
)

func main() {
	cnt := byte(0)

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		println("Can't open the port!")
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		defer conn.Close()
		if err != nil {
			println("Connection refused!")
		} else {
			var b = []byte{cnt}
			cnt++
			conn.Write(b)
		}
	}
}
