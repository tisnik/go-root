// Seriál "Programovací jazyk Go"
//
// Čtvrtá část
//
// Demonstrační příklad číslo 19:
//    Jednoduchá implementace workerů.

package main

import (
	"fmt"
	"time"
)

func worker(id int, task_channel chan int, worker_done chan bool) {
	fmt.Printf("worker %d started\n", id)
	for {
		value, more := <-task_channel
		if more {
			fmt.Printf("worker %d received task with parameter %d\n", id, value)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Printf("finishing worker %d\n", id)
			worker_done <- true
			fmt.Printf("worker %d finished\n", id)
			return
		}
	}
}

func main() {
	task_channel := make(chan int)
	worker_done := make(chan bool)

	fmt.Println("main begin")

	for i := 1; i <= 3; i++ {
		go worker(i, task_channel, worker_done)
	}
	time.Sleep(2 * time.Second)

	for i := 1; i <= 10; i++ {
		fmt.Printf("sending task with parameter %d\n", i)
		task_channel <- i
	}
	close(task_channel)

	fmt.Println("waiting for workers...")

	code, status := <-worker_done

	fmt.Printf("received code: %t and status: %t\n", code, status)
	fmt.Println("main end")
}
