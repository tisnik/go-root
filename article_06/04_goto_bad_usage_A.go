// Seriál "Programovací jazyk Go"
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 4:
//    Špatné použití příkazu goto.

package main

import "fmt"

func main() {
	goto Next
	i := 10
Next:
	fmt.Printf("%2d\n", i)
}
