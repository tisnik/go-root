// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 19:
//    Všechny unární operátory jazyka Go.

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

	p_i := &i
	fmt.Println(*p_i)

	// unární operátor <-
	channel := make(chan int)
	go message(channel)
	channel <- 1
}
