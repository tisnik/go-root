// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 4:
//     Server, který dokáže obsloužit více klientů současně
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_11/04_multi_connection_server.html

package main

import (
	"net"
)

func main() {
	cnt := byte(0)

	l, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		println("Can't open the port!")
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			println("Connection refused!")
		} else {
			println("Connected")
			go func(c net.Conn) {
				var b = []byte{cnt}
				cnt++
				c.Write(b)
				c.Close()
			}(conn)
		}
	}
}
