// Seriál "Programovací jazyk Go"
//
// Čtvrtá část
//    Rozhraní, metody, gorutiny a kanály v programovacím jazyku Go
//    https://www.root.cz/clanky/rozhrani-metody-gorutiny-a-kanaly-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 18:
//    Jednoduchá implementace workerů.

package main

import "fmt"

func worker(taskChannel chan int, worker_done chan bool) {
	fmt.Println("worker started")
	for {
		value, more := <-taskChannel
		if more {
			fmt.Printf("worker received task with parameter %d\n", value)
		} else {
			fmt.Println("finishing worker")
			worker_done <- true
			fmt.Println("worker finished")
			return
		}
	}
}

func main() {
	taskChannel := make(chan int)
	worker_done := make(chan bool)

	fmt.Println("main begin")

	go worker(taskChannel, worker_done)

	for i := 1; i <= 10; i++ {
		fmt.Printf("sending task with parameter %d\n", i)
		taskChannel <- i
	}
	close(taskChannel)

	fmt.Println("waiting for workers...")

	code, status := <-worker_done

	fmt.Printf("received code: %t and status: %t\n", code, status)
	fmt.Println("main end")
}
