// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
//
// Seznam příkladů ze sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_07/README.md
//
// Demonstrační příklad číslo 9:
//    Použití příkazu select pro zápis do kanálu.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_07/09B_select_statement_send.html

package main

import (
	"fmt"
	"time"
)

func worker(channel chan int, worker int) {
	for true {
		value, ok := <-channel
		if ok {
			fmt.Printf("Worker %d přijal hodnotu %d\n", worker, value)
		} else {
			fmt.Printf("Kanál je uzavřen pro workera %d\n", worker)
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch1 := make(chan int, 20)

	go worker(ch1, 1)
	go worker(ch1, 2)

	for i := 0; i < 10; i++ {
		select {
		case ch1 <- 0:
			fmt.Println("Poslána nula")
		case ch1 <- 1:
			fmt.Println("Poslána jednička")
		}
	}
}
