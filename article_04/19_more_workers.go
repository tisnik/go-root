// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtvrtá část
//    Rozhraní, metody, gorutiny a kanály v programovacím jazyku Go
//    https://www.root.cz/clanky/rozhrani-metody-gorutiny-a-kanaly-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 19:
//    Jednoduchá implementace workerů.

package main

import (
	"fmt"
	"time"
)

func worker(id int, taskChannel chan int, workerDone chan bool) {
	fmt.Printf("worker %d started\n", id)
	for {
		value, more := <-taskChannel
		if more {
			fmt.Printf("worker %d received task with parameter %d\n", id, value)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Printf("finishing worker %d\n", id)
			workerDone <- true
			fmt.Printf("worker %d finished\n", id)
			return
		}
	}
}

func main() {
	taskChannel := make(chan int)
	workerDone := make(chan bool)

	fmt.Println("main begin")

	for i := 1; i <= 3; i++ {
		go worker(i, taskChannel, workerDone)
	}
	time.Sleep(2 * time.Second)

	for i := 1; i <= 10; i++ {
		fmt.Printf("sending task with parameter %d\n", i)
		taskChannel <- i
	}
	close(taskChannel)

	fmt.Println("waiting for workers...")

	code, status := <-workerDone

	fmt.Printf("received code: %t and status: %t\n", code, status)
	fmt.Println("main end")
}
