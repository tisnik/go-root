package main

import (
	"log"
	"net"
	"time"
)

func processRequest(connection net.Conn, cnt *byte) {
	defer func() {
		log.Println("Closing connection")
		err := connection.Close()
		if err != nil {
			log.Println("Closing connection failed", err)
		}
	}()

	var buffer = []byte{*cnt}
	*cnt++
	n, err := connection.Write(buffer)
	if err != nil {
		log.Println("Writing error", err)
	} else {
		log.Printf("Written %d byte(s)", n)
	}
}

func main() {
	cnt := byte(0)

	l, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Can't open the port!")
	}
	defer func() {
		err := l.Close()
		if err != nil {
			log.Println("Listener close failed", err)
		}
	}()

	for {
		connection, err := l.Accept()

		if err != nil {
			log.Println("Connection refused!")
		} else {
			log.Println("Connection accepted")
			time.Sleep(2 * time.Second)
			processRequest(connection, &cnt)
		}
	}
}
