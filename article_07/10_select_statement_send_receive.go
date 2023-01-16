// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_07/README.md
//
// Demonstrační příklad číslo 10:
//    Použití příkazu select pro zápis do kanálu i pro čtení z kanálu.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_07/10_select_statement_send_receive.html

package main

import (
	"fmt"
	"time"
)

func receiver(channel chan int, receiver int) {
	for true {
		value, ok := <-channel
		if ok {
			fmt.Printf("Příjemce %d přijal hodnotu %d\n", receiver, value)
		} else {
			fmt.Printf("Kanál je pro příjemce %d uzavřen\n", receiver)
		}
		time.Sleep(2 * time.Second)
	}
}

func sender(channel chan int, sender int) {
	fmt.Printf("Odesílatel %d byl spuštěn\n", sender)
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		channel <- i
	}
	fmt.Printf("Odesílatel %d byl ukončen\n", sender)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	// ch1 := make(chan int, 20)

	go receiver(ch1, 1)
	go receiver(ch1, 2)
	go sender(ch2, 1)
	go sender(ch3, 2)

	for i := 0; i < 20; i++ {
		select {
		case ch1 <- 0:
			fmt.Println("Poslána nula")
		case ch1 <- 1:
			fmt.Println("Poslána jednička")
		case data, ok := <-ch2:
			if ok {
				fmt.Printf("Přijata data %d z kanálu 2\n", data)
			}
		case data, ok := <-ch3:
			if ok {
				fmt.Printf("Přijata data %d z kanálu 3\n", data)
			}
		}
	}
}
