// Seriál "Programovací jazyk Go"
//
// Jedenáctá část
//
// Demonstrační příklad číslo 1:
//    Klient, který přečte ze serveru sekvenci bajtů

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
