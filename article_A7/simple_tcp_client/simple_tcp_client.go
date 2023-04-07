package main

import (
	"log"
	"net"
)

const BufferSize = 1024

func main() {
	connection, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connection refused!")
		return
	}
	log.Print("Connection established, waiting for data")

	defer func() {
		log.Println("Closing connection")
		err := connection.Close()
		if err != nil {
			log.Println("Closing connection failed", err)
		}
	}()

	buffer := make([]byte, BufferSize)
	n, err := connection.Read(buffer)
	if err != nil {
		log.Println("No response!", err)
	} else {
		if n == 1 {
			log.Printf("Received %d byte: %v\n", n, buffer[:n])
		} else {
			log.Printf("Received %d bytes: %v\n", n, buffer[:n])
		}
	}
}
