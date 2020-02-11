// Seriál "Programovací jazyk Go"
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
//
// Demonstrační příklad číslo 6:
//    Použití příkazu select pro čtení z kanálů.

package main

import (
	"fmt"
	"time"
)

func worker(channel chan int, worker int) {
	fmt.Printf("Worker %d spuštěn\n", worker)
	time.Sleep(2 * time.Second)
	channel <- 1
	fmt.Printf("Worker %d ukončen\n", worker)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go worker(ch1, 1)
	go worker(ch2, 2)

	select {
	case <-ch1:
		fmt.Println("Data z kanálu 1")
	case <-ch2:
		fmt.Println("Data z kanálu 2")
	}
}
