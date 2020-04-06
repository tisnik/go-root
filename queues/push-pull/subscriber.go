package main

import (
	"github.com/go-stomp/stomp"
)

const serverAddr = "localhost:61613"
const queueName = "/queue/go_test"

func receiveMessages() {
	conn, err := stomp.Dial("tcp", serverAddr, nil)

	if err != nil {
		println("cannot connect to server", err.Error())
		return
	}
	println("connected to server", serverAddr)

	defer conn.Disconnect()

	sub, err := conn.Subscribe(queueName, stomp.AckAuto)
	if err != nil {
		println("cannot subscribe to", queueName, err.Error())
		return
	}
	for {
		msg := <-sub.C
		text := string(msg.Body)
		if text != "EXIT" {
			println("Received message", text)
		} else {
			println("Received EXIT command")
			break
		}
	}
	println("receiver finished")
}

func main() {
	receiveMessages()
}
