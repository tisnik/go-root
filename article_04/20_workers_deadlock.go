// Seriál "Programovací jazyk Go"
//
// Čtvrtá část
//    Rozhraní, metody, gorutiny a kanály v programovacím jazyku Go
//    https://www.root.cz/clanky/rozhrani-metody-gorutiny-a-kanaly-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 20:
//    Deadlock.

package main

import "fmt"

func worker(taskChannel chan int, workerDone chan bool) {
	fmt.Println("worker started")
	for {
		value, more := <-taskChannel
		if more {
			fmt.Printf("worker received task with parameter %d\n", value)
		} else {
			fmt.Println("finishing worker")
			workerDone <- true
			fmt.Println("worker finished")
			return
		}
	}
}

func main() {
	taskChannel := make(chan int)
	workerDone := make(chan bool)

	fmt.Println("main begin")

	go worker(taskChannel, workerDone)

	for i := 1; i <= 10; i++ {
		fmt.Printf("sending task with parameter %d\n", i)
		taskChannel <- i
	}
	// !!!
	// close(taskChannel)
	// !!!

	fmt.Println("waiting for workers...")

	code, status := <-workerDone

	fmt.Printf("received code: %t and status: %t\n", code, status)
	fmt.Println("main end")
}
