// Seriál "Programovací jazyk Go"
//
// Jedenáctá část
//
// Demonstrační příklad číslo 6C:
//     Připojení je nutné ukončit v gorutině, ne mimo ni

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
