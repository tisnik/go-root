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

func worker(task_channel chan int, worker_done chan bool) {
	fmt.Println("worker started")
	for {
		value, more := <-task_channel
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
	task_channel := make(chan int)
	worker_done := make(chan bool)

	fmt.Println("main begin")

	go worker(task_channel, worker_done)

	for i := 1; i <= 10; i++ {
		fmt.Printf("sending task with parameter %d\n", i)
		task_channel <- i
	}
	// !!!
	// close(task_channel)
	// !!!

	fmt.Println("waiting for workers...")

	code, status := <-worker_done

	fmt.Printf("received code: %t and status: %t\n", code, status)
	fmt.Println("main end")
}
