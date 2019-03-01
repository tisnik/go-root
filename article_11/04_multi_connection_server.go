// Seriál "Programovací jazyk Go"
//
// Jedenáctá část
//
// Demonstrační příklad číslo 4:
//     Server, který dokáže obsloužit více klientů současně

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
