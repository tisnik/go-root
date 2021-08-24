// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 1:
//    Použití příkazu goto.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_06/01_goto.html

package main

import "fmt"

func main() {
	i := 10
Next_i:
	fmt.Printf("%2d\n", i)
	i--
	if i >= 0 {
		goto Next_i
	}
}
