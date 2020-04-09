package main

import (
	"fmt"
	"github.com/go-stomp/stomp"
	"time"
)

const serverAddr = "localhost:61613"
const queueName = "/queue/go_test"

const messageCount = 10

var stop = make(chan bool)

func sendMessages() {
	defer func() {
		stop <- true
	}()

	conn, err := stomp.Dial("tcp", serverAddr, nil)
	if err != nil {
		println("Cannot connect to server", err.Error())
		return
	}
	println("Publisher part connected to server", serverAddr)
	defer conn.Disconnect()

	time.Sleep(5 * time.Second)

	for i := 1; i <= messageCount; i++ {
		text := fmt.Sprintf("Message #%d", i)
		err = conn.Send(queueName, "text/plain", []byte(text), nil)
		if err != nil {
			println("Failed to send to server", err)
			return
		}
		println("Message sent")
	}
	println("Sending EXIT message")
	err = conn.Send(queueName, "text/plain", []byte("EXIT"), nil)
	if err != nil {
		println("Failed to send EXIT message to server", err)
		return
	}
	println("Message sent")
	println("Publisher finished")
}

func receiveMessages(subscribed chan bool) {
	defer func() {
		stop <- true
	}()

	conn, err := stomp.Dial("tcp", serverAddr, nil)

	if err != nil {
		println("Cannot connect to server", err.Error())
		return
	} else {
		println("Subscriber part connected to server", serverAddr)
	}

	defer conn.Disconnect()

	time.Sleep(5 * time.Second)

	sub, err := conn.Subscribe(queueName, stomp.AckAuto)
	if err != nil {
		println("Cannot subscribe to", queueName, err.Error())
		return
	}
	close(subscribed)

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
	println("Subscriber finished")
}

func main() {
	go sendMessages()
	subscribed := make(chan bool)
	go receiveMessages(subscribed)

	<-subscribed

	<-stop
	<-stop
}
