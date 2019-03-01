// Seriál "Programovací jazyk Go"
//
// Jedenáctá část
//
// Demonstrační příklad číslo 2B:
//     Úprava adresy v předchozím příkladu

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
