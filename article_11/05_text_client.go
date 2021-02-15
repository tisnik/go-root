// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 5:
//     	Jednoduchý klient akceptující celý textový řádek

package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		println("Connection refused!")
	} else {
		fmt.Fprintf(conn, "Hello")
		status, err := bufio.NewReader(conn).ReadString('\n')
		println(status, err)
		if err != nil {
			println("No response!")
		} else {
			println(status)
		}
	}
}
