// Seriál "Programovací jazyk Go"
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 2:
//     Jednoduchý server posílající jediný bajt klientovi

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
