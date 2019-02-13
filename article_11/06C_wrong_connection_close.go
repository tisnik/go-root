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
		go func(c net.Conn) {
			defer c.Close()
			fmt.Fprintf(c, "Holla\n")
		}(conn)
	}
}
