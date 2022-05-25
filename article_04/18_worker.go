// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtvrtá část
//    Rozhraní, metody, gorutiny a kanály v programovacím jazyku Go
//    https://www.root.cz/clanky/rozhrani-metody-gorutiny-a-kanaly-v-programovacim-jazyku-go/
//
// Seznam příkladů ze čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_04/README.md
//
// Demonstrační příklad číslo 18:
//    Jednoduchá implementace workerů.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_04/18_worker.html

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
	close(taskChannel)

	fmt.Println("waiting for workers...")

	code, status := <-workerDone

	fmt.Printf("received code: %t and status: %t\n", code, status)
	fmt.Println("main end")
}
