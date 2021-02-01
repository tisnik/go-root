// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 7:
//    Použití příkazu defer.

package main

import "fmt"

func onFinish() {
	fmt.Println("Finished")
}

func main() {
	defer onFinish()

	for i := 10; i >= 0; i-- {
		fmt.Printf("%2d\n", i)
	}
	fmt.Println("Finishing main() function")
}
