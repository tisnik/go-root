// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
//
// Demonstrační příklad číslo 8:
//    Použití příkazu select pro zápis do kanálu.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_07/08_select_statement_send.html

package main

import "fmt"

func worker(channel chan int) {
	for true {
		value, ok := <-channel
		if ok {
			fmt.Printf("Přijata hodnota %d\n", value)
		} else {
			fmt.Printf("Kanál je uzavřen\n")
		}
	}
}

func main() {
	ch1 := make(chan int)

	go worker(ch1)

	for i := 0; i < 10; i++ {
		select {
		case ch1 <- 0:
			fmt.Println("Poslána nula")
		case ch1 <- 1:
			fmt.Println("Poslána jednička")
		}
	}
}
