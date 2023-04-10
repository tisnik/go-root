package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const BufferSize = 1024
const SocketFileName = "/tmp/haproxy.sock"

func reader(r io.Reader) string {
	output := ""
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil {
			log.Println("End of stream")
			return output
		}
		str := string(buf[0:n])
		output += str
		log.Println("Data read via socket", str)
	}
}

func sendCommandThroughSocket(socketFileName string, command string) (string, error) {
	connection, err := net.Dial("unix", socketFileName)
	if err != nil {
		return "", err
	}
	log.Print("Connection established, waiting for data")
	defer func() {
		log.Println("Closing connection")
		err := connection.Close()
		if err != nil {
			log.Println("Closing connection failed", err)
		}
	}()

	_, err = connection.Write([]byte(command))
	if err != nil {
		return "", err
	}
	time.Sleep(1e9)
	return reader(connection), nil
}

func main() {
	result, err := sendCommandThroughSocket(SocketFileName, "show info\n")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	fmt.Println("------------------------")
	time.Sleep(1e9)

	result, err = sendCommandThroughSocket(SocketFileName, "disable server http/myserver\n")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	fmt.Println("------------------------")
	time.Sleep(1e9)

	result, err = sendCommandThroughSocket(SocketFileName, "enable server http/myserver\n")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	fmt.Println("------------------------")
}
