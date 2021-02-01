// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 9:
//    Použití příkazu defer, zavolání funkce s parametrem.

package main

import "fmt"

func onFinish(message string) {
	fmt.Println(message)
}

func main() {
	defer onFinish("Finished")

	for i := 10; i >= 0; i-- {
		fmt.Printf("%2d\n", i)
	}
	fmt.Println("Finishing main() function")
}
