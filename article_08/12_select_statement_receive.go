// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmá část
//     Tvorba balíčků a pokročilejší operace s kanály v jazyce Go
//     https://www.root.cz/clanky/tvorba-balicku-a-pokrocilejsi-operace-s-kanaly-v-jazyce-go/
//
// Demonstrační příklad číslo 12:
//    Použití příkazu select pro čtení z kanálů s timeoutem.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_08/12_select_statement_receive.html

package main

import (
	"fmt"
	"time"
)

func worker(channel chan int, worker int) {
	fmt.Printf("Worker %d spuštěn\n", worker)
	time.Sleep(1 * time.Second)
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
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout!")
	}
}
