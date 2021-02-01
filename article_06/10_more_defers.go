// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 10:
//    Zjištění pořadí volání funkcí deklarovaných v defer().

package main

import "fmt"

func onFinish(i int) {
	fmt.Printf("Defer #%2d\n", i)
}

func main() {
	for i := 0; i <= 10; i++ {
		defer onFinish(i)
	}
	fmt.Println("Finishing main() function")
}
