// Seriál "Programovací jazyk Go"
//
// Jedenáctá část
//
// Demonstrační příklad číslo 3:
//     Server odpovídající klientovi opožděně

package main

import (
	"net"
	"time"
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
		println("connection accepted")
		time.Sleep(2 * time.Second)
		defer conn.Close()
		if err != nil {
			println("Connection refused!")
		} else {
			var b = []byte{cnt}
			cnt++
			conn.Write(b)
		}
		println("connection closed")
	}
}
