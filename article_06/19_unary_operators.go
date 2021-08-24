// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 19:
//    Všechny unární operátory jazyka Go.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_06/19_unary_operators.html

package main

import "fmt"

func message(channel chan int) {
	code, status := <-channel

	fmt.Printf("received code: %d and status: %t\n", code, status)
}

func main() {
	// unární operátory + a -
	i := 42
	fmt.Println(+i)
	fmt.Println(-i)

	// unární operátor ^
	i = 0
	fmt.Println(^i)
	i++
	fmt.Println(^i)

	// unární operátor !
	b := false
	fmt.Println(!b)

	// unární operátory & a *
	fmt.Println(&i)

	ptrI := &i
	fmt.Println(*ptrI)

	// unární operátor <-
	channel := make(chan int)
	go message(channel)
	<-channel
}
