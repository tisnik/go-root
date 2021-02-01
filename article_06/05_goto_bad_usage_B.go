// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 5:
//    Špatné použití příkazu goto.

package main

import "fmt"

func main() {
	i := 10
	goto IntoIf
	if i > 0 {
	IntoIf:
		fmt.Printf("%2d\n", i)
	}
}
